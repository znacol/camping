import { Component, OnInit, Input } from '@angular/core';
import { finalize } from 'rxjs/operators';

import { site } from '../../site';
import { ApiService } from '../../services/api.service';
import { HttpClient } from '@angular/common/http';

@Component({
    selector: 'app-create-site',
    templateUrl: './create-site.component.html',
    styleUrls: ['./create-site.component.scss'],
})
export class CreateSiteComponent implements OnInit {
    @Input() newSite: site;
    submitted = false;
    dataLoaded = true;
    nationalForests: any;
    districts: any;

    constructor(private apiService: ApiService, private http: HttpClient) {}

    ngOnInit() {
        // Fetch national forest and district info for creation dropdown
        this.apiService
            .getNationalForests()
            .subscribe(results => this.onForestsLoaded(results), err => console.log(err, 'Error loading forests'));
        this.apiService
            .getDistricts()
            .subscribe(results => this.onDistrictsLoaded(results), err => console.log(err, 'Error loading districts'));
    }

    onForestsLoaded = (results: any) => {
        this.nationalForests = results;
    };

    onDistrictsLoaded = (results: any) => {
        this.districts = results;
        this.dataLoaded = true;
    };

    submitSite = form => {
        this.submitted = true;

        // TODO refresh list of sites
        this.http
            .put('//camping.api.localhost/v1/camping/sites', {
                latitude: form.value.latitude,
                longitude: form.value.longitude,
                national_forest_id: form.value.forest,
                district_id: form.value.district,
                altitude: form.value.altitude,
                notes: form.value.notes,
            })
            .pipe(finalize(() => form.reset())) // TODO navigate to details view
            .subscribe(
                _ => { },
                err => {
                    console.log(err, 'Error creating site');
                },
            );

        // this.apiService
        //     .createSite(form.value.latitude, form.value.longitude, form.value.forest, form.value.district, form.value.altitude, form.value.notes)
        //     .subscribe(_ => {}, err => console.log(err, 'Error creating site'));
    };
}
