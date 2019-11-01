import { Component, OnInit, Input } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { finalize } from 'rxjs/operators';

import { site } from '../../site';

@Component({
  selector: 'app-create-site',
  templateUrl: './create-site.component.html',
  styleUrls: ['./create-site.component.scss']
})
export class CreateSiteComponent implements OnInit {
  @Input() newSite: site;
  submitted = false;
  dataLoaded = true;
  nationalForests: any;
  districts: any;

  constructor(private http: HttpClient) { }

  ngOnInit() {
    // Fetch national forest and district info for creation dropdown
    this.http.get('//camping.api.localhost/v1/camping/forests', {})
        .pipe(finalize(() => this.dataLoaded = true))
        .subscribe(
          results => {
            this.nationalForests = results;
          },
          err => {
            console.log(err, 'Failed to retrieve national forests');
          }
        );

    this.http.get('//camping.api.localhost/v1/camping/districts', {})
        .pipe(finalize(() => this.dataLoaded = true))
        .subscribe(
          results => {
            this.districts = results;
          },
          err => {
            console.log(err, 'Failed to retrieve districts');
          }
        );
  }

  submitSite = (form) => {
    this.submitted = true;

    // TODO refresh list of sites
    this.http.post('//camping.api.localhost/v1/camping/site', {latitude: form.value.latitude, longitude: form.value.longitude, national_forest_id: form.value.forest, district_id: form.value.district, altitude: form.value.altitude, notes: form.value.notes})
        .pipe(finalize(() => form.reset())) // TODO navigate to details view
        .subscribe(
          results => {
            console.log(results)
          },
          err => {
            console.log(err, 'Error creating site');
          }
        );
  }
}
