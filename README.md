# Catalog

Catalog is a service for product catalogs.

## Onboarding and Development Guide

### Dependencies

- Git
- Docker
- Go 1.17
- Golang Migrate (For database migration. See [golang-migrate installation](https://github.com/golang-migrate/migrate))
- Swag (For generating API docs with Swagger. See [swag installation](https://github.com/swaggo/swag))

### Installation

### Getting started
1. Clone this repository and install the prerequisites above
2. Copy `.env` from `env.sample` and modify the configuration value appropriately
3. Run the container `make compose`
4. Setup the database `make create-db && make migrate-up`
5. Download modules `go mod download`
5. Run the service by running `make start`

### Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
