# All Things Camping

### Requirements
* go
* Docker
* MySQL 8.0
* Docker Compose
* protoc
* npm

### Local Dev
* npm install in `web` directory

### Containers
* api
* database
* angular app
* swagger docs

### Scripts
* `compile-proto.sh` generates all necessary proto files and swagger docs
* `compodoc -p tsconfig.json -r 4600 -s` angular documentation


### Ports
* swagger docs: `4500`
* database: `5302`
* gRPC: `30251`
* HTTP: `8081`
* web: `4200`
* compodoc: `4600`