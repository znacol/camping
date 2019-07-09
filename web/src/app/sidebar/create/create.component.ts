import { Component, OnInit, Input } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { finalize } from 'rxjs/operators';

import { Site } from '../../Site'

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.scss']
})
export class CreateComponent implements OnInit {
  @Input() newSite: Site;
  submitted = false;

  constructor(private http: HttpClient) { }

  // TODO make national forest + district dropdown
  ngOnInit() {
  }

  submitSite(form){
    this.submitted = true;

    this.http.post('http://localhost:8000/v1/camping/site', {latitude: form.value.latitude, longitude: form.value.longitude, national_forest_id: form.value.national_forest_id, district_id: form.value.district_id, altitude: form.value.altitude, notes: form.value.notes})
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
