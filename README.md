# Go AlGo
[![Build Status](https://travis-ci.org/anaskhan96/Go-AlGo.svg?branch=master)](https://travis-ci.org/anaskhan96/Go-AlGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/anaskhan96/Go-AlGo)](https://goreportcard.com/report/github.com/anaskhan96/Go-AlGo)

This repository contains the algorithms (implemented in [Go](https://golang.org/)) taught during *Design and Analysis of Algorithms Laboratory* (CS255) course in *PES University*, Bangalore. All files have been grouped according to the week in which they were implemented.

## Testing

Navigate into any particular algorithm's folder and run `go test` command to check if it's working correctly.

You can also run `go test ./...` from the main directory itself to perform tests on all of them. Include `-v` for verbose.

Example :
```bash
$ cd Week1/SequentialSearch
$ go test
```

## Execution

Navigate into any particular algorithm's folder and run `go run file_name.go` command.

Example :
```bash
$ cd Week1/SequentialSearch
$ go run seqsearch.go
```
