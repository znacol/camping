# All Things Camping

### Requirements
* go
* Docker
* MySQL 8.0
* Docker Compose
* protoc
* npm

### Local Dev
* `docker-compose up -d --build` to build containers
* In the `web` directory, run `ng serve` to serve the web
  * `localhost:4200`

### Containers
* api
* database
* angular app
  * `localhost:4200`
* swagger (API) docs
  * Navigate to `localhost:4500` to view

### Scripts
* `compile-proto.sh` generates all necessary proto files and swagger docs
* `compodoc -p tsconfig.json -r 4600 -s` angular documentation
  * Navigate to `localhost:4600` to view


### Ports
* swagger docs: `4500`
* database: `5302`
* gRPC: `30251`
* HTTP: `8081`
* web: `4200`
* compodoc: `4600`