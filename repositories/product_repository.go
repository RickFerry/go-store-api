package repositories

import "store-api/models"

var products []models.Product

func GetAllProducts() []models.Product {
    return products
}

func GetProductByID(id uint) (models.Product, bool) {
    for _, product := range products {
        if product.ID == id {
            return product, true
        }
    }
    return models.Product{}, false
}

func CreateProduct(product models.Product) models.Product {
    products = append(products, product)
    return product
}

func UpdateProduct(id uint, updatedProduct models.Product) (models.Product, bool) {
    for i, product := range products {
        if product.ID == id {
            products[i] = updatedProduct
            return updatedProduct, true
        }
    }
    return models.Product{}, false
}

func DeleteProduct(id uint) bool {
    for i, product := range products {
        if product.ID == id {
            products = append(products[:i], products[i+1:]...)
            return true
        }
    }
    return false
}
