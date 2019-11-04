import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { ApiService } from './services/api.service';

import { site } from './site';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {
    // TODO handle multiple map types
    mapType = 'roadmap';
    sites: site[] = [];
    selectedSite: site;
    newSite: site;

    constructor(private apiService: ApiService, private http: HttpClient) {}

    ngOnInit(): void {
        // Fetch all sites
        // TODO create API service
        this.http.get('//camping.api.localhost/v1/camping/sites', {}).subscribe(
            results => {
                this.onSitesLoaded(results);
            },
            err => {
                console.log(err, 'Error occurred');
            },
        );

        // this.apiService
        //   .getAllSites()
        //   .subscribe(
        //     results => this.onSitesLoaded(results),
        //     err => console.log(err, 'Error loading sites')
        //   );
        //
    }

    onSitesLoaded = (results: any) => {
        for (const result of results.sites) {
            this.sites.push(result);
        }
    };

    markerSelected = (id: number) => {
        // Get site info from ID
        this.selectedSite = this.sites.find(i => i.id === id);
        // TODO fix logic...
        this.newSite = undefined;
    };

    siteClicked = event => {
        this.newSite = event.coords;
        // TODO fix logic...
        this.selectedSite = undefined;
    };
}
