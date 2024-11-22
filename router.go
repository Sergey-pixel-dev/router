package router

import (
	"fmt"
	"strings"
	//"fmt"
	"net/http"
)

type Router struct {
	NotFoundHandler         http.Handler
	MethodNotAllowedHandler http.Handler
	routes                  []*Route
	//namedRoutes             map[string]*Route
	//middlewears?
}

type Route struct {
	handler http.Handler
	path    string
	//name    string
	methods []string
}

/* type CHandler struct {
	http.Handler
} */

func NewRouter() *Router {
	return &Router{}
}

// NewRoute ("GET, POST, XXAXXAX", "/count/add/pi", hanldermy)
func NewRoute(methods string, path string, handler func(http.ResponseWriter, *http.Request)) *Route {
	methods = strings.Replace(methods, " ", "", 2000)
	return &Route{
		handler: http.HandlerFunc(handler),
		path:    path,
		methods: strings.Split(methods, ","),
	}
}

func (r *Router) AddRoute(route *Route) {
	r.routes = append(r.routes, route)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	PathWasFound := false
	fmt.Println(path)
	for _, route := range router.routes {
		if route.path == path {
			for _, route := range router.routes {
				if IsEqualPath(route.path, path) {
					if Contains(route.methods, r.Method) {
						route.handler.ServeHTTP(w, r)
						return
					}
					PathWasFound = true
				}
			}
		}
	}
	if PathWasFound {
		//fmt.Println("path was found but not method")
		router.MethodNotAllowedHandler.ServeHTTP(w, r)
		return
	}
	//fmt.Println("not found handler")
	router.NotFoundHandler.ServeHTTP(w, r)
}
