# Seckill_srv Service

This is the Seckill_srv service

Generated with

```
micro new seckill_srv --namespace=xlf.seckill --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: xlf.seckill.srv.seckill_srv
- Type: srv
- Alias: seckill_srv

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
./seckill_srv-srv
```

Build a docker image
```
make docker
```