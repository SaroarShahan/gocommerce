package database

var Products []Product

type Product struct {
	ID    int `json:"id"`
	Title  string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImageUrl string `json:"imageUrl"`
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

	Products = []Product{product1, product2, product3}
}