package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"strings"
	. "testTools/src/matrix/errors"
	. "testTools/src/utils/http_helpers"
)

type ProxyHttp struct {
	Method     string
	Host       string
	Path       string
	Query      map[string]string
	Protocol   string
	Header     map[string]string
	Body       map[string]string
}


func Proxy(req *http.Request) (data interface{}, errorType int, message string) {
	httpRequestIn := new(ProxyHttp)
	//httpRequestOut := new(ProxyHttp)

	err := UnmarshalHttpBody(req, httpRequestIn)
	if err != nil {
		return nil, InvalidInput, MsgUnmarshalRequestFailed
	}

	requestHost := httpRequestIn.Host
	requestPath := httpRequestIn.Path
	requestPath := httpRequestIn.Path
	requestPath := httpRequestIn.Path
	data := url.Values{}
	data.Set("name", "xiaohua")
	data.Set("id", "654321")

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	fmt.Println("[client.Do] request2 sent successfully.")

	return httpRequestIn, 200, "bbbbbbbbbbbbbbbbbbbbbbbbb"
}