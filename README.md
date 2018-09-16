![Ellcrys Network](https://storage.googleapis.com/ellcrys-docs/ellcrys-github-banner.png)

# Elld - Official Ellcrys Client
[![GoDoc](https://godoc.org/github.com/ellcrys/elld?status.svg)](https://godoc.org/github.com/ellcrys/elld)
[![CircleCI](https://circleci.com/gh/ellcrys/elld/tree/master.svg?style=svg)](https://circleci.com/gh/ellcrys/elld/tree/master)
[![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/ellnet)
[![Go Report Card](https://goreportcard.com/badge/github.com/ellcrys/elld)](https://goreportcard.com/report/github.com/ellcrys/elld)

Elld is the official client that implements a full node according to the Ellcrys specification. The client is written in Go programming language.

This client is still very much under active development. It will allow users run a daemon that will follow the protocol of the Ellcrys network. It will connect to other nodes on the network, receive and relay transactions and other messages, maintain the ledger and more.

Find more documentations in the [Wiki](https://github.com/ellcrys/elld/wiki) and in specific package directories. 

### Requirement
[Go](http://golang.org/) 1.9 or newer.

### Contributing
We use [Dep](https://github.com/golang/dep) tool to manage project dependencies. You will need it to create deterministic builds with other developers.

#### Get the Dep
Checkout the Dep [documentation](https://github.com/golang/dep#installation) for installation guide.

#### Get the source and build
```
git clone https://github.com/ellcrys/elld $GOPATH/src/github.com/ellcrys/elld
cd $GOPATH/src/github.com/ellcrys/elld
dep ensure
go install
```

## Contact
- Email: hello@ellcrys.co
- [Telegram](https://t.me/ellcryshq)
- [Twitter](https://twitter.com/ellcryshq)
