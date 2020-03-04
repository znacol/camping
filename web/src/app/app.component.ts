import { Component, OnInit, ViewChild } from '@angular/core';
import { ApiService } from './services/api.service';
import { GoogleMapsModule } from '@angular/google-maps';
import { MapInfoWindow, MapMarker } from '@angular/google-maps';

import { site } from './site';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {
    @ViewChild(MapInfoWindow, {static: false}) infoWindow: MapInfoWindow;

    center = {lat: 24, lng: 12};
    markerOptions = {draggable: false};
    markerPositions: google.maps.LatLngLiteral[] = [];
    zoom = 4;
    display?: google.maps.LatLngLiteral;

    ngOnInit(): void {
        throw new Error('Method not implemented.');
    }

    addMarker(event: google.maps.MouseEvent) {
        this.markerPositions.push(event.latLng.toJSON());
    }

    move(event: google.maps.MouseEvent) {
        this.display = event.latLng.toJSON();
    }

    openInfoWindow(marker: MapMarker) {
        this.infoWindow.open(marker);
    }

    removeLastMarker() {
        this.markerPositions.pop();
    }

    // // TODO: handle multiple map types
    // mapType = 'roadmap';
    // sites: site[] = [];
    // selectedSite: site;
    // newSite: site;
    //
    // constructor(private apiService: ApiService) {}
    //
    // ngOnInit(): void {
    //     // Fetch all sites
    //     this.apiService.getAllSites().subscribe(results => this.onSitesLoaded(results), err => console.log(err, 'Error loading sites'));
    // }
    //
    // onSitesLoaded = (results: any) => {
    //     for (const result of results.sites) {
    //         this.sites.push(result);
    //     }
    // };
    //
    // markerSelected = (id: number) => {
    //     // Get site info from ID
    //     this.selectedSite = this.sites.find(i => i.id === id);
    //     // TODO: fix logic...
    //     this.newSite = undefined;
    // };
    //
    // siteClicked = event => {
    //     this.newSite = event.coords;
    //     // TODO: fix logic...
    //     this.selectedSite = undefined;
    // };
}
