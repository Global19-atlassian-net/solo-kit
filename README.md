# solo-kit
A collection of code generation and libraries to for API development.

### Description:
- Define your declarative API in `.proto` files
- APIs are defined by top-level protobuf messages in `.proto` files
- Run Solo Kit's code generation and import generated files as libraries into your application. 
- These libraries provide an opinionated suite of tools for building a stateless, event-driven application.
- Currently only Go is supported, but other languages may be supported in the future.
- We are looking for community feedback and support!

### Examples
See `test/mock_resources.proto` and `test/generate.go` for an example of how to use solo-kit

## Build
- clone repo to gopath 
- gather dependencies: `dep ensure -v`
- use binary produced with `make solo-kit-gen` or import `cmd.Run(...)` into your own code gen code 

## Usage
- re-run whenever you change or add an api (.proto file)
- api objects generated from messages defined in protobuf files with magic comments prefixed with `@solo-kit`
- run `solo-kit-gen` recursively at the root of an `api` directory containing one or more `solo-kit.json` files
- generated files have the `.sk.go` suffix (generated test files do not include this suffix)

## Developer Setup

Solo-Kit has the following requirements:

- `git`
- `go`
- `dep`
- `protoc` (`solo-io` projects are built using version `3.6.1`)
- the `github.com/gogo/protobuf` go package

To install all the requirements, run the following:

On macOS:

```bash
# install protoc
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-osx-x86_64.zip
unzip protoc-3.6.1-osx-x86_64.zip
sudo mv bin/protoc /usr/local/bin/
rm -rf bin include protoc-3.6.1-osx-x86_64.zip readme.txt

# install go
curl https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash

```

On linux:
```bash
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
unzip protoc-3.6.1-linux-x86_64.zip
sudo mv bin/protoc /usr/local/bin/
rm -rf bin include protoc-3.6.1-linux-x86_64.zip readme.txt

# install go
curl https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash

# install dep

```
