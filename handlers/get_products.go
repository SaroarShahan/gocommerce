package handlers

import (
	"gocommerce/database"
	"gocommerce/utils"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	utils.SendData(w, database.Products, http.StatusOK)
}