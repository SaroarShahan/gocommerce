package handlers

import (
	"encoding/json"
	"gocommerce/database"
	"gocommerce/utils"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(database.Products) + 1;
	database.Products = append(database.Products, newProduct)

	utils.SendData(w, newProduct, http.StatusCreated)
}