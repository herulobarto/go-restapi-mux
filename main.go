package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/herulobarto/go-restapi-mux/controllers/productcontroller"
	"github.com/herulobarto/go-restapi-mux/models"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/products", productcontroller.Index).Methods("GET")
	r.HandleFunc("/product/{id}", productcontroller.Show).Methods("GET")
	r.HandleFunc("/product", productcontroller.Create).Methods("POST")
	r.HandleFunc("/product/{id}", productcontroller.Update).Methods("PUT")
	r.HandleFunc("/product", productcontroller.Delete).Methods("DELETE")

	http.ListenAndServe(":8080", r)

}
