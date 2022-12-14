# Cat Facts API

Simple cat facts api built with golang and mongodb.
Feel free to steal it.

### Features

The api has only one route which fetches all cat facts (`json` list format):

`GET http://localhost:3000/facts`

There is a worker that accesses `https://catfact.ninja/fact` api asynchronously and populates a mongodb collection, which is accessed by the user.


## Setup

### .env
See `.env-example`for reference
### mongodb
run `docker-compose up -d` with `docker` or install mongo locally.
> Make sure the .env matches your docker config

## Run
Run `make run` if you have `make` installed, or else `go run .`
### Parameters
All parameters can be changed from the `main.go` file
## Build
Run `make build` or `go build`
