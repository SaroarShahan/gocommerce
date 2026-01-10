package cmd

import (
	"fmt"
	"gocommerce/global_router"
	"gocommerce/middlewares"
	"net/http"
)

func Serve() {
	manager := middlewares.MiddlewareManager()
	manager.Use(middlewares.Logger)
	mux := http.NewServeMux()

	initRoutes(mux, manager)

	fmt.Println("Server started on port 8080")

	router := global_router.HandleGlobalRouter(mux)
	
	err := http.ListenAndServe(":8080", router)
	
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}