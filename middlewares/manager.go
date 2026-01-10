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

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
		n := next

		// for i := len(middlewares) - 1; i >= 0; i-- {
		// 	middleware := middlewares[i]
		// 	n = middleware(n)
		// }

		for _, middleware := range middlewares {
			n = middleware(n)
		}

		return n
}