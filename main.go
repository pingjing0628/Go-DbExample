package main

import(
	"GoMySQL/apis/product_api"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/product/findall", product_api.FindAll).Methods("GET")
	router.HandleFunc("/api/product/search/{keyword}", product_api.Search).Methods("GET")
	router.HandleFunc("/api/product/searchprices/{min}/{max}", product_api.SearchPrices).Methods("GET")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}