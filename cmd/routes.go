package cmd

import (
	"gocommerce/handlers"
	"gocommerce/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProduct)))
}