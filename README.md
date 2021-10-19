# `bowditch`

API to fetch [Cal Dining menus](https://caldining.berkeley.edu/menus/) in JSON format.

## Routes
`GET /:restaurant` - fetches all meals for the restaurant specified

`GET /:restaurant/:meal` - fetches specified meal for the restaurant specified

## Running locally
```go
$ go mod download
$ go build
$ ./bowditch
```
