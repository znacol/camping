import { Component } from '@angular/core';
import { Request } from "../../../api/proto/api_pb.js";
import { CampingServiceClient } from "../../../api/proto/api_grpc_web_pb.js";


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

var campingService = new CampingServiceClient('https://localhost:8080');

const req = new Request();
req.setMessage("test");

campingService.do(req, {}, function(err, response) {
  console.log("called api")
  console.log(err)
  console.log(response)
});

// grpc.invoke(CampingService.Do, {
//   request: req,
//   host: "127.0.0.1:8080",
//   // onMessage: (message: Message) => {
//   //   console.log("got message", message.toObject());
//   // },
//   onEnd: (code: grpc.Code, msg: string | undefined, trailers: BrowserHeaders) => {
//     if (code == grpc.Code.OK) {
//       console.log("all ok")
//     } else {
//       console.log("hit an error", code, msg, trailers);
//     }
//   }
// });