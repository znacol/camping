import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

export class AppComponent {
  latitude = -28.68352;
  longitude = -147.20785;
  mapType = 'satellite';
  title = 'camping';
}