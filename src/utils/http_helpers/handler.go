package http_helpers

import (
	"net/http"
	"runtime/debug"
	. "testTools/src/utils/clog"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"testTools/src/utils/metric"
)

const (
	LabelUrl = "url"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

// rateMetric must only have label Url
func Wrap(rateMetric metric.Meter, f HandlerFunc) HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		if rateMetric != nil {
			rateMetric.Inc(prometheus.Labels{LabelUrl: req.URL.RequestURI()})
		}
		// print request trace info
		PrintRequestInfo(req)
		start := time.Now()
		defer func() {
			r := recover()
			if r != nil && r != http.ErrAbortHandler {
				Hlog.SetRequest(req).
					SetTag(LF_ResponseTime, time.Since(start).String()).
					Errorf("Recover in Request, stack: %s %v", string(debug.Stack()), r)
			}
		}()
		Hlog.SetRequest(req).
			Info("Request start")
		f(w, req)
		Hlog.SetRequest(req).
			SetTag(LF_ResponseTime, time.Since(start).String()).
			Info("Request done")
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

type ThickHandlerFunc func(w http.ResponseWriter, req *http.Request) (
	data interface{}, /*api-specified response result*/
	err int,
	message string /*reason if err not empty*/)

func WrapThick(f ThickHandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// call handler
		data, errorType, msg := f(w, req)
		// make response
		res := GetResponse(req, data, errorType, msg)
		// do response
		DoResponse(w, res)
	}
}

func AddHandleFunc(r *mux.Router, method string, path string, f HandlerFunc) {
	r.Methods(method).Path(path).HandlerFunc(Wrap(nil, f))
}

func AddThinHandleFunc(r *mux.Router, method string, path string, f ThinHandlerFunc) {
	AddHandleFunc(r, method, path, WrapTHF(f))
}

func AddThickHandleFunc(r *mux.Router, method string, path string, f ThickHandlerFunc) {
	AddHandleFunc(r, method, path, WrapThick(f))
}

func DoWrapTHF(w http.ResponseWriter, req *http.Request,
	data interface{}, errorType int, msg string) {
	if data == nil {
		data = ""
	}
	res := new(Response)
	res.Code = errorType
	if errorType == NoError {
		res.Data = data
	} else {
		res.Message = msg
		Hlog.SetRequest(req).
			Debugf("Request error(%v) message(%s)", errorType, msg)
	}
	DoResponse(w, res)
}
