# User_srv Service

This is the User_srv service

Generated with

```
micro new user_srv --namespace=xlf.user --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: xlf.user.srv.user_srv
- Type: srv
- Alias: user_srv

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
./user_srv-srv
```

Build a docker image
```
make docker
```