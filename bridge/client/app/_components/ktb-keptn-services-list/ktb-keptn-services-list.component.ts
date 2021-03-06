import {
  ChangeDetectorRef,
  Component,
  EventEmitter,
  Input, OnDestroy,
  OnInit,
  Output
} from '@angular/core';
import {DtTableDataSource} from '@dynatrace/barista-components/table';
import {UniformRegistration} from "../../_models/uniform-registration";
import {BehaviorSubject, Observable, Subject} from "rxjs";
import {DataService} from "../../_services/data.service";
import {ActivatedRoute} from "@angular/router";
import {switchMap, takeUntil} from "rxjs/operators";
import {UniformRegistrationLog} from "../../_models/uniform-registration-log";

@Component({
  selector: 'ktb-keptn-services-list',
  templateUrl: './ktb-keptn-services-list.component.html',
  styleUrls: ['./ktb-keptn-services-list.component.scss']
})
export class KtbKeptnServicesListComponent implements OnInit, OnDestroy {

  private readonly unsubscribe$ = new Subject<void>();

  public uniformRegistrations: DtTableDataSource<UniformRegistration> = new DtTableDataSource();

  private selectedUniformRegistrationId$ = new Subject<string>();
  private uniformRegistrationLogsSubject = new BehaviorSubject<UniformRegistrationLog[]>([]);

  public selectedUniformRegistration: UniformRegistration;
  public uniformRegistrationLogs$: Observable<UniformRegistrationLog[]> = this.uniformRegistrationLogsSubject.asObservable();
  public isLoadingLogs = false;

  public projectName: string;

  @Output() selectedUniformRegistrationChanged: EventEmitter<UniformRegistration> = new EventEmitter();

  constructor(private dataService: DataService, private route: ActivatedRoute) {
  }

  ngOnInit(): void {
    this.route.paramMap.pipe(
      takeUntil(this.unsubscribe$)
    ).subscribe(map => {
      this.projectName = map.get('projectName');
    });

    this.selectedUniformRegistrationId$.pipe(
      takeUntil(this.unsubscribe$),
      switchMap(uniformRegistrationId => {
        this.isLoadingLogs = true;
        return this.dataService.getUniformRegistrationLogs(uniformRegistrationId);
      })
    ).subscribe((uniformRegLogs) => {
      uniformRegLogs.sort((a, b) => {
        if (a.time.valueOf() > b.time.valueOf()) return -1;
        if (a.time.valueOf() < b.time.valueOf()) return 1;
        return 0;
      });
      this.isLoadingLogs = false;
      this.uniformRegistrationLogsSubject.next(uniformRegLogs);
    });

    this.dataService.getUniformRegistrations()
      .subscribe((uniformRegistrations) => {
        this.uniformRegistrations.data = uniformRegistrations;
      });
  }

  ngOnDestroy(): void {
    this.unsubscribe$.next();
  }

  public setSelectedUniformRegistration(uniformRegistration: UniformRegistration) {
    if (this.selectedUniformRegistration !== uniformRegistration) {
      this.selectedUniformRegistration = uniformRegistration;
      this.selectedUniformRegistrationId$.next(this.selectedUniformRegistration.id);
      this.selectedUniformRegistrationChanged.emit(uniformRegistration);
    }
  }

  public formatSubscriptions(subscriptions: string[]): string {
    return subscriptions.join('<br/>');
  }

  public sortData(sortEvent) {
    if(this.uniformRegistrations.data) {
      const isAscending = sortEvent.direction === 'asc';
      const data: UniformRegistration[] = this.uniformRegistrations.data.slice();

      data.sort((a, b) => {
        switch (sortEvent.active) {
          case 'name': return this.compare(a.name, b.name, isAscending);
          case 'host': return (this.compare(a.metadata.hostname, b.metadata.hostname, isAscending) || this.compare(a.name, b.name, true));
          case 'namespace': return this.compare(a.metadata.kubernetesmetadata.namespace, b.metadata.kubernetesmetadata.namespace, isAscending) || this.compare(a.name, b.name, true);
          case 'location': return this.compare(a.metadata.location, b.metadata.location, isAscending) || this.compare(a.name, b.name, true);
        }
      });

      this.uniformRegistrations.data = data;
    } else {
      this.uniformRegistrations.data = [];
    }
  }

  private compare(a: string, b: string, isAsc: boolean): number {
    const result = a.localeCompare(b);
    if (result !== 0 && !isAsc) {
      return -result;
    }
    return result;
  }
}
