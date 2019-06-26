# All Things Camping

### Requirements
* go
* Docker
* MySQL 8.0
* Docker Compose
* protoc
* npm
    * `npm install @angular/cli`
    * `npm i @agm/core --save` (run in `frontend` directory)
    * `npm install google-protobuf @types/google-protobuf grpc-web-client ts-protoc-gen --save`


### Proto
* Compiling proto: `protoc -I . api.proto --go_out=plugins=grpc:.`
* `protoc -I=. api.proto  --js_out=import_style=commonjs:.  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.`