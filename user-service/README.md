# User Service

This is the User service

Generated with

```
micro new user-service --namespace=ke.micro.book --alias=user --type=service
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: ke.micro.book.service.user
- Type: service
- Alias: user

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./user-service
```

Build a docker image
```
make docker
```