package router

import (
	"github.com/gorilla/mux"
	"testTools/src/matrix/handler"
)

func NewRouter() *mux.Router{
	//init router
	r := mux.NewRouter()
	r.UseEncodedPath()

	// rule helper
	add := func(method string, path string, f handler.HandlerFunc) {
		r.Methods(method).Path(path).HandlerFunc(f)
	}
	//addTHF := func(method string, path string, f ThinHandlerFunc) {
	//	add(method, path, WrapTHF(f))
	//}

	add("GET", "/", handler.HelloWorldHandler)
	return r
}
