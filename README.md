# All Things Camping

**THIS IS STILL IN DEVELOPMENT AND IN A VERY ROUGH STATE.**

Web app that utilizes postgres, google maps, golang, and angular to save, explore, and map camp sites.

### Requirements
* Docker Compose

### Local Dev (on OSX)
* `docker-compose up -d` to start containers
    * Web changes will compile automatically on save
    * API changes require running `docker-compose restart camping-api`
 * Navigate to `camping.app.localhost` to view app

### Scripts
* `./scripts/compile-proto.sh` generates all necessary proto files and swagger docs
* `compodoc -p tsconfig.json -r 4600 -s` generates and serves angular documentation
  * Navigate to `localhost:4600` to view
