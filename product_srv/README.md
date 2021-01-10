# Product_srv Service

This is the Product_srv service

Generated with

```
micro new product_srv --namespace=xlf.product --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: xlf.product.srv.product_srv
- Type: srv
- Alias: product_srv

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
./product_srv-srv
```

Build a docker image
```
make docker
```