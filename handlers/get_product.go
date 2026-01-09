package handlers

import (
	"gocommerce/database"
	"gocommerce/utils"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	id, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range database.Products {
		if product.ID == id {
			utils.SendData(w, product, http.StatusOK)
			return
		}
	}

	utils.SendData(w, "Product not found", http.StatusNotFound)
}