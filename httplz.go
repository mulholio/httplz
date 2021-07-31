package main

import (
	"fmt"
)

type codeNumber = int16
type statusCode struct {
	Code              codeNumber
	Name, Description string
}

func (sc statusCode) Format() string {
	return fmt.Sprintf("\n%d: %s\n\n%s", sc.Code, sc.Name, sc.Description)
}

var codes = map[codeNumber]statusCode{
	400: statusCode{Code: 400, Name: "Bad Request", Description: "The HyperText Transfer Protocol (HTTP) 400 Bad Request response status code indicates that the server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing)."},
}

func main() {
	// TODO flags to read
	c := codes[400]
	fmt.Println(c.Format())
}
