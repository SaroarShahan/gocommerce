package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID    int `json:"id"`
	Title  string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImageUrl string `json:"imageUrl"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Us Page")
}

var products []Product

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	handleCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(products)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	handleCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(products) + 1;
	products = append(products, newProduct)

	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
}

func main() {
	router := http.NewServeMux()

	router.Handle("GET /", http.HandlerFunc(homeHandler))

	router.Handle("GET /about", http.HandlerFunc(aboutHandler))

	router.Handle("GET /products", http.HandlerFunc(getProductsHandler))
	router.Handle("POST /create-products", http.HandlerFunc(createProductHandler))

	fmt.Println("Server started on port 8080")
	
	err := http.ListenAndServe(":8080", router)
	
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func init() {
	product1 := Product{
		ID:          1,
		Title:       "Product 1",
		Description: "Description for Product 1",
		Price:       19.99,
		ImageUrl:    "http://example.com/product1.jpg",
	}

	product2 := Product{
		ID:          2,
		Title:       "Product 2",
		Description: "Description for Product 2",
		Price:       29.99,
		ImageUrl:    "http://example.com/product2.jpg",
	}

	product3 := Product{
		ID:          3,
		Title:       "Product 3",
		Description: "Description for Product 3",
		Price:       39.99,
		ImageUrl:    "http://example.com/product3.jpg",
	}

	products = []Product{product1, product2, product3}
}