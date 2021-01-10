# Web Service

This is the Web service

Generated with

```
micro new web --namespace=xlf.web --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: xlf.web.web.web
- Type: web
- Alias: web

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./web-web
```

Build a docker image
```
make docker
```