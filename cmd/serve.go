package cmd

import (
	"fmt"
	"gocommerce/global_router"
	"gocommerce/handlers"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProduct))

	fmt.Println("Server started on port 8080")

	router := global_router.HandleGlobalRouter(mux)
	
	err := http.ListenAndServe(":8080", router)
	
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}