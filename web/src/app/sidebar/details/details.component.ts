import { Component, OnInit, OnChanges, Input } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { finalize } from 'rxjs/operators';

@Component({
  selector: 'app-details',
  templateUrl: './details.component.html',
  styleUrls: ['./details.component.scss']
})
export class DetailsComponent implements OnInit, OnChanges {
  @Input() selectedSite: any;
  nationalForest: any;
  district: any;
  dataLoaded = false;

  constructor(private http: HttpClient) { }

  ngOnInit() {
    // Fetch national forest and district info
    this.http.get('camping-api:8000/v1/camping/forest/' + this.selectedSite.national_forest_id, {})
        .subscribe(
          results => {
            this.nationalForest = results;
          },
          err => {
            console.log('Error occured');
          }
        );

    this.http.get('camping-api:8000/v1/camping/district/' + this.selectedSite.district_id, {})
        // .pipe(finalize(() => this.dataLoaded = true))
        .subscribe(
          results => {
            this.district = results;
          },
          err => {
            console.log('Error occurred');
          }
        );

    this.dataLoaded = true;
  }

  ngOnChanges() {
  }
}
