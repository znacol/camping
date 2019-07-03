import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent {
  mapType = 'roadmap';
  sites: any[] = [];
  selectedSite: any;

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    // Fetch all sites
    this.http.get('http://localhost:8081/v1/camping/sites', {
      })
        .subscribe(
          results => {
            this.onSitesLoaded(results);
          },
          err => {
            console.log('Error occured');
          }
        );
  }

  markerSelected(id: number): void {
    // Get site info from ID
    this.selectedSite = this.sites.find(i => i.id === id);

  }

  public onSitesLoaded = (results) => {
    for (const result of results.sites) {
      this.sites.push(result);
    }
  }

}
