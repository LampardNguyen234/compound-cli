[![Go Report Card](https://goreportcard.com/badge/github.com/LampardNguyen234/compound-cli)](https://goreportcard.com/report/github.com/LampardNguyen234/compound-cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/LampardNguyen234/compound-cli/blob/main/LICENSE)

compound-cli
=============
A simple CLI for doing things that are beyond the capabilities of the regular SDK.

<!-- toc -->
* [Usage](#usage)
* [Commands](./commands.md)
<!-- tocstop -->

# Usage
<!-- usage -->
## Installation
Install to the `$GOPATH` folder.
```shell
$ go install
```
This command will install the CLI application into your `GOPATH` folder. Alternatively, you can build and install the binary file
into a desired folder by the following command.
```shell
$ go build -o PATH/TO/YOUR/FOLDER/appName
```
If you have issues with these commands, try to clean the golang module cache first.
```shell
go clean --modcache
```

## Usage
See [Commands](./commands.md)