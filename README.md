# All Things Camping

### Requirements
* go
* Docker
* MySQL 8.0
* Docker Compose
* protoc
* npm
    * `npm i @angular/cli`
    * `npm i google-protobuf grpc-web @agm/core --save`


### Proto
<!-- * Compiling proto: `protoc -I . api.proto --go_out=plugins=grpc:.` -->
<!-- * `protoc -I=. api.proto  --js_out=import_style=commonjs:.  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.` -->
<!-- * `protoc -I=. api.proto --js_out=import_style=commonjs:. --grpc-web_out=import_style=typescript,mode=grpcwebtext:.` -->

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  path/to/your_service.proto


  protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  path/to/your_service.proto