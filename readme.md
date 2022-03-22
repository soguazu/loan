# Company Service

Company service manages everything that has to do with a company.

## Built with:
- [Golang](https://go.dev/dl/)
- [Postgres](https://postgresapp.com) product database
- [Sqlite](https://www.sqlite.org/download.html) product testing
- [Docker](https://www.docker.com/products/docker-desktop/)

## Installation

Use the package manager [go modules](https://go.dev/blog/using-go-modules) to install all dependencies.

```bash
git clone https://github.com/Evea01/company-service.git
```

```bash
cd company-service
```

```bash
go mod download
```

```bash
touch .env
```
Copy the value inside .env.sample into the .env and fill the values for the necessary config


## Usage

```bash
# To build a postgres docker image
make postgres

# To create a postgres database 
make createdb

# To run build swagger docs and run service
make swagger

# To run both unit and integration test
make test
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)