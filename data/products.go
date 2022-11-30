package data

import (
	"fmt"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`                          // Unique identifier for the product
	Name        string  `json:"name" validate:"required"`    // the name for this poduct
	Description string  `json:"description" `                // the description for this poduct
	Price       float32 `json:"price" validate:"gt=0"`       // the price for the product
	SKU         string  `json:"sku" validate:"required,sku"` // the SKU for the product
}

var ErrProductNotFound = fmt.Errorf("Product not found")

// Products is a collection of Product
type Products []*Product

// AddProduct adds a new product to the database
func AddProduct(p Product) {
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, &p)
}

// UpdateProduct replaces a product in the database with the given item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func UpdateProduct(p Product) error {
	i := findIndexByProductID(p.ID)
	if i == -1 {
		return ErrProductNotFound
	}
	// update the product in the DB
	productList[i] = &p
	return nil
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil
}

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

// GetProductByID returns a single product which matches the id from the database.
// If a product is not found this function returns a ProductNotFound error
func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}
	return productList[i], nil
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

// productList is a hard coded list of products for this
// example data source
var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
	},
	{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
	},
}
