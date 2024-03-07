package router

import (
	"net/http"
	"strings"
)

type MiddlewareFunc func(http.Handler) http.Handler

type middleware interface {
	Middleware(handler http.Handler) http.Handler
}

type Router struct {
	mux         *http.ServeMux
	middlewares []middleware
}

func New() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (mw MiddlewareFunc) Middleware(handler http.Handler) http.Handler {
	return mw(handler)
}

func (r *Router) Use(mwf ...MiddlewareFunc) {
	for _, fn := range mwf {
		r.middlewares = append(r.middlewares, fn)
	}
}

func (r *Router) UseWithPrefix(prefix string, mwf ...MiddlewareFunc) {
	for _, fn := range mwf {
		mw := MiddlewareFunc(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if strings.Contains(req.URL.Path, prefix) {
					fn(next).ServeHTTP(w, req)
				} else {
					next.ServeHTTP(w, req)
				}
			})
		})
		r.middlewares = append(r.middlewares, mw)
	}
}

func (r *Router) Handle(pattern string, handler http.Handler) {
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i].Middleware(handler)
	}
	r.mux.Handle(pattern, handler)
}

func (r *Router) GET(pattern string, fn http.HandlerFunc) {
	r.Handle(http.MethodGet+" "+pattern, fn)
}

func (r *Router) POST(pattern string, fn http.HandlerFunc) {
	r.Handle(http.MethodPost+" "+pattern, fn)
}

func (r *Router) DELETE(pattern string, fn http.HandlerFunc) {
	r.Handle(http.MethodDelete+" "+pattern, fn)
}

func (r *Router) PUT(pattern string, fn http.HandlerFunc) {
	r.Handle(http.MethodPut+" "+pattern, fn)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
