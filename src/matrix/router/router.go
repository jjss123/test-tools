package router

import (
	"github.com/gorilla/mux"
	"testTools/src/matrix/handler"
	. "testTools/src/utils/http_helpers"
)

func NewRouter() *mux.Router{
	//init router
	r := mux.NewRouter()
	r.UseEncodedPath()

	// rule helper
	add := func(method string, path string, f HandlerFunc) {
		r.Methods(method).Path(path).HandlerFunc(Wrap(handler.MatrixRequestRate, f))
	}
	addTHF := func(method string, path string, f ThinHandlerFunc) {
		add(method, path, WrapTHF(f))
	}

	addTHF("GET", "/", handler.Index)

	addTHF("POST", "/tool/proxy", handler.Proxy)
	return r
}
