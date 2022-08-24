# go-practice

A very simple REST API built using [go](https://go.dev/) and [gin](https://gin-gonic.com/) as a way to learn and become familiar with these technologies.

## Running the application
The project can be launched using the `make up` command. This will start up mysql in a docker container and initialize the database, then launch the webserver. To shut everything back down run `make down`.

To reload changes to the webserver without touching the database container run `make bounce`

## API Documentation
The API endpoints are documented via swagger docs which can be accessed at http://localhost:8080/docs/index.html

To re-generate the swagger documentation after making changes, run `make update-swagger` and then bounce the server to see the doc changes.