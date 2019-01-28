[![GitHub release](https://img.shields.io/github/release/igarciaolaizola/h2-forward.svg)](https://github.com/igarciaolaizola/h2-forward/releases)
[![Build Status](https://travis-ci.com/igarciaolaizola/h2-forward.svg?branch=master)](https://travis-ci.com/igarciaolaizola/h2-forward)
[![Go Report Card](https://goreportcard.com/badge/igarciaolaizola/h2-forward)](http://goreportcard.com/report/igarciaolaizola/h2-forward)
[![license](https://img.shields.io/github/license/igarciaolaizola/h2-forward.svg)](https://github.com/igarciaolaizola/h2-forward/blob/master/LICENSE.md)

# h2-forward

HTTP2 reverse proxy implemented in golang

## Usage

```
go run cmd/h2-forward/main.go --addr=localhost:8080 --port=8081
```