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

  constructor(private http: HttpClient) { }

  ngOnInit(): void {
    // Fetch all sites
    this.http.get('http://localhost:8081/v1/camping/sites', {
      })
        .subscribe(
          res => {
            for (const result of res.sites) {
              this.sites.push(result);
            }
          },
          err => {
            console.log('Error occured');
          }
        );
  }

  // TODO display sidebar
  markerSelected(id: int): void {
    console.log(id)
  }
}
