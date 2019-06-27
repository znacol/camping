import { Component } from '@angular/core';
// import { CampingServiceClient } from "./ApiServiceClientPb";
// import { Request } from "./api_pb";
import { HttpClient } from '@angular/common/http';



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

  constructor(private http: HttpClient){ }
  ngOnInit(): void {
    this.http.post('http://localhost:8081/v1/camping/do', {
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

// var campingService = new CampingServiceClient('http://localhost:8081', null, null);

// const req = new Request();
// req.setMessage("test")

// campingService.do(req, {}, function(err, response) {
//   console.log("called api")
//   console.log(err)
//   console.log(response)
// });
