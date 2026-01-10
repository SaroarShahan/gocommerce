package cmd

import (
	"fmt"
	"net/http"
	"strconv"

	"gocommerce/config"
	"gocommerce/middlewares"
)

func Serve() {
	config := config.GetConfig()
	manager := middlewares.MiddlewareManager()
	manager.Use(middlewares.Preflight, middlewares.Cors, middlewares.Logger)
	
	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)
	
	addr := ":" + strconv.Itoa(config.HttpPort)
	fmt.Println("Server started on port", config.HttpPort)
	
	err := http.ListenAndServe(addr, wrappedMux)
	
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}