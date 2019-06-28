import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})

export class AppComponent {
  latitude = -28.68352;
  longitude = -147.20785;
  mapType = 'satellite';
  title = 'camping';

  constructor(private http: HttpClient){ }
  ngOnInit(): void {
    this.http.get('http://localhost:8081/v1/camping/sites', {
      })
        .subscribe(
          res => {
            console.log(res);
          },
          err => {
            console.log("Error occured");
          }
        );
  }
}
