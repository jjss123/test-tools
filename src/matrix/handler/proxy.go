package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	. "testTools/src/matrix/errors"
	. "testTools/src/utils/http_helpers"
)

type ProxyHttpElement struct {
	Method     string
	Host       string
	Path       string
	Query      map[string]string
	Protocol   string
	Header     map[string]string
	Body       map[string]string
}

func ProxyHttp(req *http.Request) (data interface{}, errorType int, message string) {
	httpRequestIn := new(ProxyHttpElement)

	err := UnmarshalHttpBody(req, httpRequestIn)
	if err != nil {
		return nil, InvalidInput, MsgProxyUnmarshalRequestFailed
	}

	//setup method
	requestOutMethod:= httpRequestIn.Method

	//setup url
	newUrl, err := url.Parse(httpRequestIn.Host)
	if err != nil {
		return nil, InvalidInput, MsgProxyParseUrlFailed
	}
	newUrl.Path = httpRequestIn.Path
	newQuery := newUrl.Query()
	for k, v := range httpRequestIn.Query{
		newQuery.Set(k, v)
	}
	newUrl.RawQuery = newQuery.Encode()
	var requestOutUrl = newUrl.String()

	//setup body
	newBody, err := json.Marshal(httpRequestIn.Body)
	var requestOutBody = &bytes.Buffer{}
	requestOutBody.Write(newBody)

	//setup request client
	client := &http.Client{}
	newRequest, _ := http.NewRequest(requestOutMethod, requestOutUrl, requestOutBody)

	//setup header
	for k, v := range httpRequestIn.Header{
		newRequest.Header.Add(k, v)
	}
	newRequest.Header.Add("Content-Type", "application/json")

	//send request
	resp, err := client.Do(newRequest)
	if err != nil {
		fmt.Println(err.Error())
		return nil, InternalError, err.Error()
	}
	defer resp.Body.Close()

	//return
	b, _ := ioutil.ReadAll(resp.Body)
	var ret map[string]interface{}
	if err := json.Unmarshal(b, &ret); err != nil {
		if strings.Contains(err.Error(), "invalid character"){
			return string(b), NoError, NoMessage
		}
		return nil, InvalidInput, MsgProxyParseReturnBodyFailed
	}
	return ret, NoError, NoMessage
}