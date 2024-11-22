package router

import (
	"net/http"
	"strings"
)

type Router struct {
	NotFoundHandler         http.Handler
	MethodNotAllowedHandler http.Handler
	routes                  []*Route
	//namedRoutes             map[string]*Route
}

type Route struct {
	handler http.Handler
	path    string
	//name    string
	methods []string
}

/* type CustomHandler struct {
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

func (route *Route) AddMiddleware(f func(w http.ResponseWriter, r *http.Request)) *Route {
	tmp := route.handler
	route.handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f(w, r)
		tmp.ServeHTTP(w, r)
	})
	return route
}
func (router *Router) AddMiddleware(f func(w http.ResponseWriter, r *http.Request)) {
	for _, route := range router.routes {
		route.AddMiddleware(f)
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	PathWasFound := false
	for _, route := range router.routes {
		if IsEqualPath(route.path, path) {
			if Contains(route.methods, r.Method) {
				route.handler.ServeHTTP(w, r)
				return
			}
			PathWasFound = true
		}
	}
	if PathWasFound {
		router.MethodNotAllowedHandler.ServeHTTP(w, r)
		return
	}
	router.NotFoundHandler.ServeHTTP(w, r)
}
