package utils

import "net/http"


/**
 * Unified response struct for http request
 * Will be marshaled to json and put in http body
 */
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"` // api-specified response result
}

type HandlerFunc func(http.ResponseWriter, *http.Request)


// make handler func more easy to be tested
type ThinHandlerFunc func(*http.Request) (
	data interface{}, /*api-specified response result*/
	err int,
	message string /*reason if err not empty*/)

func WrapTHF(f ThinHandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// call handler
		data, errorType, msg := f(req)
		// make response
		res := GetResponse(req, data, errorType, msg)
		// do response
		DoResponse(w, res)
	}
}


func GetResponse(req *http.Request, data interface{}, errorType int, message string) *Response {
	res := new(Response)
	res.Code = errorType
	if errorType == NoError {
		res.Data = data
		res.Message = message
	} else {
		res.Message = message
		Hlog.SetRequest(req).
			Debugf("Request error(%v) message(%s)", errorType, message)
	}

	return res
}

