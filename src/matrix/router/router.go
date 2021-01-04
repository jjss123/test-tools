package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)


//func Index(req *http.Request) (data interface{}, errorType int, message string) {
//	fmt.Println("Hello World!")
//	return "Hello World!", 0, "Hello World!"
//}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request)  {
	_, _ = fmt.Fprintf(w, "HelloWorld!")
}

func NewRouter() *mux.Router{
	//init router
	r := mux.NewRouter()
	r.UseEncodedPath()

	// rule helper
	add := func(method string, path string, f HandlerFunc) {
		r.Methods(method).Path(path).HandlerFunc(f)
	}
	//addTHF := func(method string, path string, f ThinHandlerFunc) {
	//	add(method, path, WrapTHF(f))
	//}

	add("GET", "/", HelloWorldHandler)
	return r
}
