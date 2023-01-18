# navigationsvc

## What is this?

The service helps drones to locate the nearest databank via a JSON API inorder to upload gathered data from space exploration.

## Local Development

* `make test` to run unit tests (or from inside your IDE)
* `make ci` to run additional lint checks on the top of tests

## API

* Swagger Documentation is available at /doc/api folder.
* For easy access, open https://editor.swagger.io/ in the browser and paste the contents from api.yaml.

## Running service locally

* `make run` to run on default port 5055 and sector ID is 1

## Running service in Docker

* `make image` to build an image
* docker run -p 5058:5055 --env PORT=5055 navigationsvc

## Example

```
Health check 
Request
curl --location --request POST 'http://localhost:5055/'

Response
{}

------------------------------------------------------------------
Get location 
Request
curl --location --request POST 'http://localhost:5055/location' \
--data-raw '{
"x": "123.12",
"y": "456.56",
"z": "789.89",
"vel": "12.9"
}'

Response
{
    "loc": 1382.47
}
```