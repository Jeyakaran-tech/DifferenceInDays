# DifferenceInDays
[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

Made with :heart: from Jeyakaran to WeMoney

This repo contains source code to find the difference between two dates(in days). This logic is done without using any kind of third party libraries or "time" package from GoLang. 

This is a simple CLI `(Command Line Interface)` application, which can accept only dates between `01/01/1900` and `31/12/2999`. 

The format needs to be only `dd/mm/YYYY`.

# PreRequisites

- Go - `go version go1.17.6 darwin/amd64`
- Ginkgo - `Ginkgo Version 2.1.3`

# Execution

- `make run` - To start the program.
- `make unit_test` - To run all the unit tests in the project

## To install Ginkgo
```sh
go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo
go get github.com/onsi/gomega/...
```

