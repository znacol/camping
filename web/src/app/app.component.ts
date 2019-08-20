import {
  Component,
  OnInit
 } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { site } from './site';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent implements OnInit {
  // TODO handle multiple map types
  mapType = 'roadmap';
  sites: site[] = [];
  selectedSite: site;
  newSite: site;

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    // Fetch all sites
    // TODO create API service
    this.http.get('http://localhost:8000/v1/camping/sites', {
      })
        .subscribe(
          results => {
            this.onSitesLoaded(results);
          },
          err => {
            console.log(err, 'Error occurred');
          }
        );
  }

  public onSitesLoaded = (results: any) => {
    for (const result of results.sites) {
      this.sites.push(result);
    }
  }

  public markerSelected = (id: number) => {
    // Get site info from ID
    this.selectedSite = this.sites.find(i => i.id === id);
    // TODO fix logic...
    this.newSite = undefined;
  }

  public siteClicked = (event) => {
    this.newSite = event.coords;
    // TODO fix logic...
    this.selectedSite = undefined;
  }

}
