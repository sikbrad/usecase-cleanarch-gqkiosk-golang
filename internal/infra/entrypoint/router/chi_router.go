package router

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

type chiRouter struct{

}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(url string, f func(resp http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Get(url, f)
}

func (*chiRouter) POST(url string, f func(resp http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Post(url, f)
}

func (*chiRouter) SERVE(port string) {
	fmt.Println("Chi http server listening on port :",port)
	http.ListenAndServe(port, chiDispatcher)
}