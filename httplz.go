package main

import (
	"fmt"
	"os"
	"strconv"
)

type codeNumber = int64
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
	code := os.Args[1]

	if code == "" {
		panic("You must pass a code to httplz in order to receive a result.")
	}

	c, err := strconv.ParseInt(code, 10, 16)
	if err != nil {
		panic("Status code passed must be an int")
	}

	statusCode, ok := codes[c]
	if !ok {
		panic("Could not find status code with that value")
	}

	fmt.Println(statusCode.Format())
}
