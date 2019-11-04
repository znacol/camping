import { Component, OnInit, OnChanges, Input } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { finalize } from 'rxjs/operators';

@Component({
    selector: 'app-details',
    templateUrl: './details.component.html',
    styleUrls: ['./details.component.scss'],
})
export class DetailsComponent implements OnInit, OnChanges {
    @Input() selectedSite: any;
    nationalForest: any;
    district: any;
    dataLoaded = false;

    constructor(private http: HttpClient) {}

    ngOnInit() {
        // Fetch national forest and district info
        this.http.get('//camping.api.localhost/v1/camping/forests/' + this.selectedSite.national_forest_id, {}).subscribe(
            results => {
                this.onForestsLoaded(results);
            },
            err => {
                console.log(err, 'Error occurred');
            },
        );

        this.http.get('//camping.api.localhost/v1/camping/districts/' + this.selectedSite.district_id, {}).subscribe(
            results => {
                this.onDistrictsLoaded(results);
            },
            err => {
                console.log(err, 'Error occurred');
            },
        );
    }

    onForestsLoaded = (results: any) => {
        this.nationalForest = results.forests[0];
    };

    onDistrictsLoaded = (results: any) => {
        this.district = results.districts[0];
        this.dataLoaded = true;
    };

    ngOnChanges() {}
}
