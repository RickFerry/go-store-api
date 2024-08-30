package controllers

import (
    "encoding/json"
    "net/http"
    "store-api/models"
    "store-api/services"
    "strconv"

    "github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
    products := services.GetAllProducts()
    json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    product, found := services.GetProductByID(uint(id))
    if found {
        json.NewEncoder(w).Encode(product)
    } else {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode("Product not found")
    }
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    _ = json.NewDecoder(r.Body).Decode(&product)
    createdProduct := services.CreateProduct(product)
    json.NewEncoder(w).Encode(createdProduct)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var updatedProduct models.Product
    _ = json.NewDecoder(r.Body).Decode(&updatedProduct)
    product, updated := services.UpdateProduct(uint(id), updatedProduct)
    if updated {
        json.NewEncoder(w).Encode(product)
    } else {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode("Product not found")
    }
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    deleted := services.DeleteProduct(uint(id))
    if deleted {
        json.NewEncoder(w).Encode("Product deleted")
    } else {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode("Product not found")
    }
}
