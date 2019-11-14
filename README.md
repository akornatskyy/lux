# lux

[![Build Status](https://travis-ci.org/akornatskyy/lux.svg?branch=master)](https://travis-ci.org/akornatskyy/lux) [![Go Report Card](https://goreportcard.com/badge/github.com/akornatskyy/lux)](https://goreportcard.com/report/github.com/akornatskyy/lux) [![Go Doc](https://godoc.org/github.com/akornatskyy/lux?status.svg)](https://godoc.org/github.com/akornatskyy/lux)

A package manager for shell scripts.

## Install

```sh
go get github.com/akornatskyy/lux
```

or download a binary from available releases.

## Usage

```sh
usage: lux [options...] [ns/]package[:bundle][ ...]
  -u    update package
  -v    verbose mode
packages @ https://raw.githubusercontent.com/akornatskyy/lux-pkg
```

Example:

```sh
lux lua:5.1
```

## Packages

Feel free to contribute packages to [central](https://github.com/akornatskyy/lux-pkg) repository by submitting a pull request or create a repository named *lux-pkg* in your github.com account. In the later case your account name would serve as a namespace for package.

You can override the central repository URL by exporting *LUX_URL* environment varilable. Packages and downloads are stored under *~/.cache/lux/*,  use *LUX_PKG* environment variable to override.

## Release

```sh
CGO_ENABLED=0 go build -ldflags "-w -s" && upx --ultra-brute lux
```
