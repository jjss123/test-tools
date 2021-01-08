package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	. "testTools/src/matrix/errors"
	. "testTools/src/utils/http_helpers"
	"testTools/src/utils/marshaler"
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

	err := UnmarshalHttpBody(req, httpRequestIn)
	if err != nil {
		return nil, InvalidInput, MsgUnmarshalRequestFailed
	}

	//setup url
	u, err := url.Parse(httpRequestIn.Host)
	if err != nil {
		return nil, InvalidInput, MsgProxyParseUrlFailed
	}
	u.Path = httpRequestIn.Path
	q := u.Query()
	for k, v := range httpRequestIn.Query{
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	var urlStr = u.String()

	//setup body
	requestOutBody, err := json.Marshal(httpRequestIn.Body)
	var body = &bytes.Buffer{}
	body.Write(requestOutBody)

	//setup request client
	client := &http.Client{}
	r, _ := http.NewRequest(httpRequestIn.Method, urlStr, body)

	//setup header
	for k, v := range httpRequestIn.Header{
		r.Header.Add(k, v)
	}
	r.Header.Add("Content-Type", "application/json")

	//send request
	resp, err := client.Do(r)
	if err != nil {
		return nil, InvalidInput, MsgProxyRequestFailed
	}
	defer resp.Body.Close()
	//return
	b, err := ioutil.ReadAll(resp.Body)

	a := new(interface{})

	a, err = json.Marshal(b)
	fmt.Print(b)
	return string(c), NoError, NoMessage
}