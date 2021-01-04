package handler

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)


func Index(req *http.Request) (data interface{}, errorType int, message string) {
	fmt.Println("Hello World!")
	return "Hello World!", 0, "Hello World!"
}