# CanoDB

[![Go Reference](https://pkg.go.dev/badge/github.com/lindsuen/canodb/ferretdb.svg)](https://pkg.go.dev/github.com/lindsuen/canodb)
[![Go Report Card](https://goreportcard.com/badge/github.com/lindsuen/canodb)](https://goreportcard.com/report/github.com/lindsuen/canodb)
[![build](https://github.com/lindsuen/canodb/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/lindsuen/canodb/actions/workflows/build.yml)
![GitHub Release](https://img.shields.io/github/v/release/lindsuen/canodb)
![GitHub License](https://img.shields.io/github/license/lindsuen/canodb)

CanoDB is a key-value database based on [goLevelDB](https://github.com/syndtr/goleveldb).

--------

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [License](#license)

--------

## Features

## Requirements

- Go 1.20 or higher.

## Installation

You should use `go get` to add the library in your repository.

```sh
$ go get github.com/lindsuen/canodb/v1
```

The compiled `canodb-cli` can be downloaded from [the release page](https://github.com/lindsuen/canodb/releases).  
Also, you can build from the source code by yourself. During the compilation process, `git` and `make` are necessary.

```sh
$ git clone https://github.com/lindsuen/canodb
$ cd canodb/
$ make build
```

## Usage

### Command-line tool `canodb-cli`

Export data

```sh
$ canodb-cli -d data/ -e canodb.dmp
```

Import data

```sh
$ canodb-cli -d data/ -i canodb.dmp
```

### How to use leveldb package

Please see [here](https://github.com/lindsuen/canodb/blob/master/docs/LEVELDB.md) for details.

## License

[BSD-2-Clause license](https://github.com/lindsuen/canodb/blob/master/LICENSE)
