# go-freebox
[![GoDoc](https://godoc.org/github.com/moul/go-freebox?status.svg)](https://godoc.org/github.com/juju2013/go-freebox)


:wrench: Freebox API v4 client in Golang, deviated sensibly from moul's version

## About

This is golang api binding for [Freebox](https://en.wikipedia.org/wiki/Freebox). It's a subset including:

 * Login and authroization request
 * AirMedia
 * Call Entries
 * Contact
 
There's also a cli application demonstrating the API.

## Install

`go get -u github.com/juju2013/go-freebox`

`GOFBX_TOKEN= go run github.com/juju2013/go-freebox/cmd/freebox`

### Usage

The API make use of the following environ variables:

 * GOFBX_ID, GOFBX_NAME: any string you like
 * GOFBX_TOKEN: Authorization token, dpends on values of ID and NAME
 * GOFBX_LOGLEVEL: can be INFO, DEBUG, WARN and ERROR, default to INFO
 
### Authorization

To use the API, you need first to get a GOFBX_TOKEN:

 * Give a value to GOFBX_ID and GOFBX_NAME, the auth token depends on them
 * Unset GOFBX_TOKEN, Call Authorize(), the Freebox device will prompt you on the small LCD display to accept/deny, accept then and the token will be found in Client.App.token (and also printed)
 * Set GOFBX_TOKEN to this token and you won't need to do the authorization again
 
 
 
