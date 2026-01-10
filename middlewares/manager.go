package middlewares

import "net/http"

type Middleware func(next http.Handler) http.Handler

type Manager struct {
	middlewares []Middleware
}

func MiddlewareManager() *Manager {
	return &Manager{
		middlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.middlewares = append(mngr.middlewares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
		n := next

		for _, middleware := range middlewares {
			n = middleware(n)
		}

		return n
}

func (mngr *Manager) WrapMux(next http.Handler) http.Handler {
		n := next

		for _, middleware := range mngr.middlewares {
			n = middleware(n)
		}

		return n
}