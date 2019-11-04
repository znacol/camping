import { Component, OnInit, OnChanges, Input } from '@angular/core';
import { ApiService } from '../../services/api.service';

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

    constructor(private apiService: ApiService) {}

    ngOnInit() {
        // Fetch national forest and district info
        this.apiService
            .getNationalForests(this.selectedSite.national_forest_id)
            .subscribe(results => this.onForestsLoaded(results), err => console.log(err, 'Error loading forests'));
        this.apiService
            .getDistricts(this.selectedSite.district_id)
            .subscribe(results => this.onDistrictsLoaded(results), err => console.log(err, 'Error loading districts'));
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
