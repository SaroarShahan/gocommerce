package cmd

import (
	"fmt"
	"net/http"

	"gocommerce/middlewares"
)

func Serve() {
	manager := middlewares.MiddlewareManager()
	manager.Use(middlewares.Preflight, middlewares.Cors, middlewares.Logger)
	
	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)
	
	fmt.Println("Server started on port 8080")
	
	err := http.ListenAndServe(":8080", wrappedMux)
	
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}