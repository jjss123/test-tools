package main

import (
	"fmt"
	"matrix/router"
	"net/http"
)

func main(){
	r := router.NewRouter()
	err := http.ListenAndServe("0.0.0.0:80", r)

	if err != nil {
		fmt.Println("服务器错误")
	}
}
