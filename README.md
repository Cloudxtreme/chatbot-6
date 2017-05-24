# Chatbot

> chatbot implemetation in Golang using Scylladb for data storage

## Getting Started

Clone the repo:

```sh
$ git clone https://github.com/jaschweder/chatbot
```

### Prerequisites

 - Docker and docker compose

### Installing

Start all containers with:

```sh
$ docker-compose up -d
```

## Running the tests

```sh
$ docker-compose exec api go test
```

## Deployment

Build the API:

```sh
$ docker-compose exec api go build
```

## Built With

* [Go](https://www.golang.org/) - The Go programming language
* [ScyllDB](https://www.scylladb.com/) - ScyllaDB No-SQL database

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/jaschweder/chatbot/tags).

## Authors

* **Jonathan A. Schweder** - *Initial work* - [Bludata](https://github.com/bludata)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
