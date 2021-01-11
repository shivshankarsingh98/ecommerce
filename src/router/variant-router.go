package router

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	model "models"
	"net/http"
	"strconv"
)


func (pcs *ProductCatalogueService) initializeVariantRoutes() {
	pcs.Router.HandleFunc("/variant", pcs.createVariant).Methods("POST")//create variant
	pcs.Router.HandleFunc("/variant/{id:[0-9]+}", pcs.getVariant).Methods("GET")//get variant
	pcs.Router.HandleFunc("/variant/{id:[0-9]+}", pcs.updateVariant).Methods("PUT")//update variant
	pcs.Router.HandleFunc("/variant/{id:[0-9]+}", pcs.deleteVariant).Methods("DELETE")//delete variant
}


func (pcs *ProductCatalogueService) getVariant(w http.ResponseWriter, r *http.Request) {
	log.Println("GET variant Request from: ",r.RemoteAddr)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid variant ID")
		return
	}

	variant := model.VariantDetails{VariantId: id}
	if err := variant.GetVariant(pcs.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Variant not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, variant)
}


func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if payload != nil {
		w.Write(response)
	}
}


func (pcs *ProductCatalogueService) createVariant(w http.ResponseWriter, r *http.Request) {
	log.Println("POST variant Request from: ",r.RemoteAddr)
	queryParam := r.URL.Query()
	variant := model.VariantDetails{}
	variant.Colour = queryParam.Get("colour")
	variant.VariantName = queryParam.Get("variant_name")
	variant.Size = queryParam.Get("size")

	if discountPriceStr := queryParam.Get("discount_price");discountPriceStr != "" {
		discountPrice, err := strconv.ParseFloat(discountPriceStr, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid discount price")
			return
		}else {
			variant.DiscountPrice = discountPrice
		}
	}

	if mrpStr := queryParam.Get("mrp");mrpStr != "" {
		mrp, err := strconv.ParseFloat(mrpStr, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid mrp price")
			return
		}else {
			variant.Mrp = mrp
		}
	}

	if productIdStr := queryParam.Get("product_id");productIdStr != "" {
		productId, err := strconv.ParseInt(productIdStr, 10,64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid product id")
			return
		}else {
			variant.ProductId = productId
		}
	}

	if variantIdStr := queryParam.Get("variant_id");variantIdStr != "" {
		variantId, err := strconv.ParseInt(variantIdStr, 10,64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid variant id")
			return
		}else {
			variant.VariantId = variantId
		}
	}

	if err := variant.CreateVariant(pcs.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, variant)
}


func (pcs *ProductCatalogueService) updateVariant(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT variant Request from: ",r.RemoteAddr)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}
	queryParam := r.URL.Query()


	variant := model.VariantDetails{VariantId: id}
	variant.Colour = queryParam.Get("colour")
	variant.VariantName = queryParam.Get("variant_name")
	variant.Size = queryParam.Get("size")

	if discountPriceStr := queryParam.Get("discount_price");discountPriceStr != "" {
		discountPrice, err := strconv.ParseFloat(discountPriceStr, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid discount price")
			return
		}else {
			variant.DiscountPrice = discountPrice
		}
	}else {
		variant.DiscountPrice = -1
	}

	if mrpStr := queryParam.Get("mrp");mrpStr != "" {
		mrp, err := strconv.ParseFloat(mrpStr, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid mrp price")
			return
		}else {
			variant.Mrp = mrp
		}
	}else {
		variant.Mrp = -1
	}

	if productIdStr := queryParam.Get("product_id");productIdStr != "" {
		productId, err := strconv.ParseInt(productIdStr, 10,64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid product id")
			return
		}else {
			variant.ProductId = productId
		}
	}else{
		variant.ProductId = -1
	}


	if err = variant.UpdateVariant(pcs.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}

func (pcs *ProductCatalogueService) deleteVariant(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE variant Request from: ",r.RemoteAddr)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid variant ID")
		return
	}

	variant := model.VariantDetails{VariantId: id}
	if err := variant.DeleteVariant(pcs.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, nil)


}
