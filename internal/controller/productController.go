package controller

import (
	"Hi_Tech/internal/model"
	"Hi_Tech/internal/repository"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	model.Product
}

type ProductController struct {
	productRepository *repository.ProductRepository
}

func NewProductController(productRepository *repository.ProductRepository) *ProductController {
	return &ProductController{productRepository: productRepository}
}

func (pc *ProductController) GetAll(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all products")
	products, _ := pc.productRepository.GetAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (pc *ProductController) Add(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid product", http.StatusBadRequest)
		return
	}

	pc.productRepository.Add(product)
	log.Printf("Inserted product: %s", product.Name)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Product inserted successfully")
}

func (pc *ProductController) Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	productID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	pc.productRepository.Remove(productID)
	log.Printf("Removed product with id: %d", productID)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Product removed successfully")
}

func main() {
	db, err := sql.Open("postgres", "user=your_user password=your_password dbname=your_db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
