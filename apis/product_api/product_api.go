package product_api

import (
	"github.com/gorilla/mux"
	"GoMySQL/config"
	"encoding/json"
	"net/http"
	"GoMySQL/models"
	"strconv"
)

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJson(response, http.StatusOK, products)
		}
	}
}

func Search(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := config.GetDB()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Search(keyword)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJson(response, http.StatusOK, products)
		}
	}
}

func SearchPrices(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	smin := vars["min"]
	smax := vars["max"]
	min, _ := strconv.ParseFloat(smin, 64)
	max, _ := strconv.ParseFloat(smax, 4)
	db, err := config.GetDB()
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.SearchPrices(min, max)
		if err2 != nil {
			responseWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			responseWithJson(response, http.StatusOK, products)
		}
	}
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	responseWithJson(w, code, map[string]string{"error": msg})
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}