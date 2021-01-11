package router

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	model "models"
	"net/http"
	"strconv"
)

func (pcs *ProductCatalogueService) initializeCategoryRoutes() {
	pcs.Router.HandleFunc("/category", pcs.createCategory).Methods("POST")//create category
	pcs.Router.HandleFunc("/category/{id:[0-9]+}", pcs.getCategory).Methods("GET")//get category
	pcs.Router.HandleFunc("/category/{id:[0-9]+}", pcs.updateCategory).Methods("PUT")//update category
	pcs.Router.HandleFunc("/category/{id:[0-9]+}", pcs.deleteCategory).Methods("DELETE")//delete category
}

func (pcs *ProductCatalogueService) getCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("GET category Request from: ",r.RemoteAddr)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid variant ID")
		return
	}

	category := model.CategoryDetails{CategoryId: id}
	if err := category.GetCategory(pcs.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Category not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, category)
}


func (pcs *ProductCatalogueService) createCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("POST category Request from: ",r.RemoteAddr)
	queryParam := r.URL.Query()
	category := model.CategoryDetails{}
	category.CategoryName = queryParam.Get("category_name")

	if categoryIdStr := queryParam.Get("category_id");categoryIdStr != "" {
		categoryId, err := strconv.ParseInt(categoryIdStr, 10,64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid product id")
			return
		}else {
			category.CategoryId = categoryId
		}
	}

	if parentCategoryIdStr := queryParam.Get("parent_category_id");parentCategoryIdStr != "" {
		parentCategoryId, err := strconv.ParseInt(parentCategoryIdStr, 10,64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid product id")
			return
		}else {
			category.ParentCategoryId = parentCategoryId
		}
	}else{
		category.ParentCategoryId = -1
	}


	if err := category.CreateCategory(pcs.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, category)
}


func (pcs *ProductCatalogueService) updateCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("PUT category Request from: ",r.RemoteAddr)

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	queryParam := r.URL.Query()
	category := model.CategoryDetails{CategoryId: id}
	category.CategoryName = queryParam.Get("category_name")

	if parentCategoryIdStr := queryParam.Get("parent_category_id");parentCategoryIdStr != "" {
		parentCategoryId, err := strconv.ParseInt(parentCategoryIdStr, 10,64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid parent category id")
			return
		}else {
			category.ParentCategoryId = parentCategoryId
		}
	}else{
		category.ParentCategoryId = -1
	}

	if err := category.UpdateCategory(pcs.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, nil)
}


func (pcs *ProductCatalogueService) deleteCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("DELETE category Request from: ",r.RemoteAddr)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid variant ID")
		return
	}

	category := model.CategoryDetails{CategoryId: id}
	if err := category.DeleteCategory(pcs.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, category)
}
