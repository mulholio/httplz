# `httplz`

A very simple HTTP status code lookup CLI tool. Built to learn rudimentary Go syntax.

## Installation

- `git clone https://github.com/mulholio/httplz`
- `go build httplz.go`

## Usage

- Pass the status code number as the sole argument:

  ```bash   
  ./httplz 400
  ```

  Output:
  ```
  400: Bad Request

  The server could not understand the request due to invalid syntax.
  ```
