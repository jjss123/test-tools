package main

import (
	"fmt"
	"net/http"
	"testTools/src/matrix/router"
)

func main() {
	r := router.NewRouter()
	err := http.ListenAndServe(":5111", r)

	if err != nil {
		fmt.Println("服务器错误")
	}
}
