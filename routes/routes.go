package routes

import (
    "store-api/controllers"
    "store-api/middleware"

    "github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()

    router.Use(middleware.LoggingMiddleware)
    
    productRouter := router.PathPrefix("/products").Subrouter()
    productRouter.Use(middleware.AuthMiddleware)

    productRouter.HandleFunc("", controllers.GetProducts).Methods("GET")
    productRouter.HandleFunc("/{id}", controllers.GetProduct).Methods("GET")
    productRouter.HandleFunc("", controllers.CreateProduct).Methods("POST")
    productRouter.HandleFunc("/{id}", controllers.UpdateProduct).Methods("PUT")
    productRouter.HandleFunc("/{id}", controllers.DeleteProduct).Methods("DELETE")

    return router
}
