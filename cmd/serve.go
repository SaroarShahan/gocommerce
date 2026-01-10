package cmd

import (
	"fmt"
	"gocommerce/global_router"
	"gocommerce/handlers"
	"gocommerce/middlewares"
	"net/http"
)

func Serve() {
	manager := middlewares.MiddlewareManager()
	mux := http.NewServeMux()

	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts), middlewares.Logger))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct), middlewares.Logger))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProduct), middlewares.Logger))

	fmt.Println("Server started on port 8080")

	router := global_router.HandleGlobalRouter(mux)
	
	err := http.ListenAndServe(":8080", router)
	
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}