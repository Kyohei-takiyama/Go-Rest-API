package router

import "net/http"


type Router interface {
	// GET is a shortcut for router.Handle("GET", path, handle)
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	// POST is a shortcut for router.Handle("POST", path, handle)
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}