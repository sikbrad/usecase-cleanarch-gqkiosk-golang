package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct {

}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router{
	return &muxRouter{}
}

func (*muxRouter) GET(url string, f func(resp http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(url,f).Methods("GET")
}

func (*muxRouter) POST(url string, f func(resp http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(url,f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Println("Mux http server listening on port :",port)
	http.ListenAndServe(port,muxDispatcher)
}