package router

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	model "models"
	"net/http"
	"strconv"
	"strings"
)

func (pcs *ProductCatalogueService) initializeProductRoutes() {
	pcs.Router.HandleFunc("/product", pcs.createProduct).Methods("POST")//create product
	pcs.Router.HandleFunc("/product/{id:[0-9]+}", pcs.getProduct).Methods("GET")//get product
	//pcs.Router.HandleFunc("/product/{id:[0-9]+}", pcs.updateProduct).Methods("PUT")//update product
	pcs.Router.HandleFunc("/product/{id:[0-9]+}", pcs.deleteProduct).Methods("DELETE")//delete product
}

func (pcs *ProductCatalogueService) getProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("GET product Request from: ",r.RemoteAddr)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid variant ID")
		return
	}

	product := model.ProductDetails{ProductId: id}
	if err := product.GetProduct(pcs.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, product)
}


func (pcs *ProductCatalogueService) createProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("POST product Request from: ",r.RemoteAddr)

	queryParam := r.URL.Query()
	product := model.ProductDetails{}
	product.ProductName = queryParam.Get("product_name")
	product.Description = queryParam.Get("description")
	product.ProductImageUrl = queryParam.Get("url")

	if productIdStr := queryParam.Get("product_id");productIdStr != "" {
		productId, err := strconv.ParseInt(productIdStr, 10,64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid product id")
			return
		}else {
			product.ProductId = productId
		}
	}

	categoryIds := queryParam.Get("category_ids")
	categoryArray := strings.Split(categoryIds, ",")
	for _, categoryIdStr := range categoryArray {
		categoryId, err := strconv.ParseInt(categoryIdStr, 10,64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid category id")
			return
		}else {
			product.CategoryIds = append(product.CategoryIds, categoryId)
		}
	}

	if err := product.CreateProduct(pcs.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, product)
}


//func (pcs *ProductCatalogueService) updateProduct(w http.ResponseWriter, r *http.Request) {
//  log.Println("PUT product Request from: ",r.RemoteAddr)

//	vars := mux.Vars(r)
//	id, err := strconv.Atoi(vars["id"])
//	if err != nil {
//		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
//		return
//	}
//
//	var p product
//	decoder := json.NewDecoder(r.Body)
//	if err := decoder.Decode(&p); err != nil {
//		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
//		return
//	}
//	defer r.Body.Close()
//	p.ID = id
//
//	if err := p.updateProduct(a.DB); err != nil {
//		respondWithError(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	respondWithJSON(w, http.StatusOK, p)
//}
//

func (pcs *ProductCatalogueService) deleteProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE product Request from: ",r.RemoteAddr)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid variant ID")
		return
	}

	product := model.ProductDetails{ProductId: id}
	if err := product.DeleteProduct(pcs.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, product)


}
