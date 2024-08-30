package services

import (
    "store-api/models"
    "store-api/repositories"
)

func GetAllProducts() []models.Product {
    return repositories.GetAllProducts()
}

func GetProductByID(id uint) (models.Product, bool) {
    return repositories.GetProductByID(id)
}

func CreateProduct(product models.Product) models.Product {
    return repositories.CreateProduct(product)
}

func UpdateProduct(id uint, updatedProduct models.Product) (models.Product, bool) {
    return repositories.UpdateProduct(id, updatedProduct)
}

func DeleteProduct(id uint) bool {
    return repositories.DeleteProduct(id)
}
