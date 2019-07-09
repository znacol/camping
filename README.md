# All Things Camping

**THIS IS STILL IN DEVELOPMENT AND IN A VERY ROUGH STATE.**

Web app that utilizes mySQL, google maps, golang gRPC, and angular to save, explore, and map camp sites.

### Requirements
* go
* Docker
* MySQL 8.0
* Docker Compose
* protoc
* npm

### Local Dev
* `docker-compose up -d` to start containers

### Containers
* api
* database
* angular app
  * Navigate to `0.0.0.0:4200`
* swagger (API) docs
  * Navigate to `localhost:4500`

### Scripts
* `compile-proto.sh` generates all necessary proto files and swagger docs
* `compodoc -p tsconfig.json -r 4600 -s` generate angular documentation and serve
  * Navigate to `localhost:4600` to view

### Ports
* swagger docs: `4500`
* database: `5302`
* gRPC: `30251`
* HTTP: `8081`
* web: `4200`
* compodoc: `4600`