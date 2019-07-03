import {
  Component,
  Input,
  OnInit,
  OnChanges,
} from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { finalize } from 'rxjs/operators';


@Component({
  selector: 'app-sidebar',
  templateUrl: 'sidebar.component.html',
  styleUrls: ['sidebar.component.scss'],
})

export class SidebarComponent implements OnInit, OnChanges {
  @Input() selectedSite: any;
  nationalForest: any;
  district: any;
  dataLoaded = false;

  constructor(private http: HttpClient) { }

  ngOnInit() { }

  ngOnChanges() {
    // Fetch national forest and district info
    this.http.get('http://api:8081/v1/camping/forest/' + this.selectedSite.national_forest_id, {})
        .subscribe(
          results => {
            this.nationalForest = results;
          },
          err => {
            console.log('Error occured');
          }
        );

    this.http.get('http://api:8081/v1/camping/district/' + this.selectedSite.district_id, {})
        .pipe(finalize(() => this.dataLoaded = true))
        .subscribe(
          results => {
            this.district = results;
          },
          err => {
            console.log('Error occured');
          }
        );
  }
}
