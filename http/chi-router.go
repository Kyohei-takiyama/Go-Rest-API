package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)


type chaiRouter struct {}

var (
	chiDispatcher = chi.NewRouter()
)


func NewChaiRouter() Router {
	return &chaiRouter{}
}


func (*chaiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri , f)
}

func (*chaiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri , f)
}

func (*chaiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server running on port : %v\n" , port)
	http.ListenAndServe(":" + port , chiDispatcher)
}