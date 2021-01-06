package handler

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)


func HelloWorldHandler(w http.ResponseWriter, r *http.Request)  {
	_, _ = fmt.Fprintf(w, "HelloWorld!")
}