package models

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

type CategoryDetails struct {
	CategoryId         int64              `json:"category_id"`
	CategoryName       string             `json:"category_name"`
	ProductsList       *[]ProductDetails   `json:"products,omitempty"`
	ParentCategoryId   int64              `json:"parent_category_id,omitempty"`
	ChildCategories    []*CategoryDetails  `json:"child_categories,omitempty"`
}

//CREATE TABLE category (
//category_id INT(6)  PRIMARY KEY,
//category_name VARCHAR(30) NOT NULL,
//parent_category_id INT(6)
//);


func GetChildProductIdsList(db *sql.DB, categoryId int64) ([]int64,error) {
	results, err := db.Query("select product_id from product_category where category_id = ?", categoryId)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var productId int64
	has_results := false
	productIdList := []int64{}
	for results.Next() {
		has_results = true
		err = results.Scan(&productId)
		if err != nil {
			return nil, err
		}
		productIdList = append(productIdList, productId)
	}
	if (!has_results) {
		return nil, sql.ErrNoRows
	}
	return productIdList, nil
}
func GetProductDetail(variantDetailsChan chan ProductDetails, productId int64,db *sql.DB){
	product := ProductDetails{ProductId: productId}
	if err := product.GetProduct(db); err != nil {
		log.Println(err)
	} else {
		variantDetailsChan <- product
	}

}
func GetChildProductsList(db *sql.DB, categoryId int64)[]ProductDetails{
	productIdList, err := GetChildProductIdsList(db,categoryId)
	if err != nil{
		return nil
	}
	productDetailsChan := make(chan ProductDetails)
	defer close(productDetailsChan)

	productDetailsList := []ProductDetails{}

	for _, productId := range productIdList {
		go GetProductDetail(productDetailsChan,productId, db)
	}

	totalProductChild := len(productIdList)
	for totalProductChild > 0 {
		productDetailsList = append(productDetailsList, <-productDetailsChan)
		totalProductChild -=1
	}
	return productDetailsList


}
func (cd *CategoryDetails) GetCategory(db *sql.DB) error {
	categoryQueue := []*CategoryDetails{}

	queryResult, err := db.Query("select category_name from category where category_id = ?", cd.CategoryId)
	if err != nil {
		return err
	}
	has_results := false
	for queryResult.Next() {
		has_results = true
		err = queryResult.Scan(&cd.CategoryName)
		if err != nil {
			return err
		}
	}
	if (!has_results) {
		return sql.ErrNoRows
	}

	categoryDetails := CategoryDetails{CategoryId: cd.CategoryId, CategoryName: cd.CategoryName}
	categoryQueue = append(categoryQueue,&categoryDetails)
	rootCategory := categoryQueue[0]

	for len(categoryQueue)>=1 {
		currentCategory := categoryQueue[0]
		results, err := db.Query("select category_id,category_name from category where parent_category_id = ?", currentCategory.CategoryId)
		if err != nil {
			return err
		}
		productDetailsList := GetChildProductsList(db,currentCategory.CategoryId)
		if len(productDetailsList) >=1 {
			currentCategory.ProductsList = &productDetailsList
		}

		var categoryId int64
		var categoryName string
		childCategoryIndex := 0
		for results.Next() {
			err = results.Scan(&categoryId, &categoryName)
			if err != nil {
				return err
			}
			childCategory := CategoryDetails{}
			childCategory.CategoryId = categoryId
			childCategory.CategoryName = categoryName
			(*currentCategory).ChildCategories = append((*currentCategory).ChildCategories, &childCategory)
			categoryQueue = append(categoryQueue,&childCategory)
			childCategoryIndex +=1
		}
		categoryQueue = categoryQueue[1:]
	}

	*cd = *rootCategory
	return nil
}

func (cd *CategoryDetails) updateCategory(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (cd *CategoryDetails) DeleteCategory(db *sql.DB) error {

	updateProductCategoryStmt, err := db.Prepare("delete from  product_category  where category_id= ?")
	if err != nil{
		return err
	}
	_, err = updateProductCategoryStmt.Exec(cd.CategoryId)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("delete from category where category_id=?")
	if err != nil{
		return err
	}
	res, err := stmt.Exec(cd.CategoryId)
	if err != nil {
		return err
	}else {
		rowEffected, err1 := res.RowsAffected()
		if err1 != nil {
			return err1
		} else if rowEffected == 0 {
			errMsg := "No category present with id: " +strconv.FormatInt(cd.CategoryId, 10)
			return errors.New(errMsg)
		}
	}

	return nil
}

func (cd *CategoryDetails) CreateCategory(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO category (category_id, category_name, parent_category_id) VALUES (?, ?, ?)")
	if err != nil{
		return err
	}
	if cd.ParentCategoryId == -1 {
		_, err = stmt.Exec(cd.CategoryId, cd.CategoryName, nil)

	}else {
		_, err = stmt.Exec(cd.CategoryId, cd.CategoryName, cd.ParentCategoryId)
	}
	if err != nil {
		return err
	}
	return nil
}


//func GetChildVariantsIdsList(db *sql.DB, productId int64) ([]int64,error) {
//	results, err := db.Query("select variant_id from variant where product_id = ?", productId)
//	if err != nil {
//		return nil, err
//	}
//	defer results.Close()
//
//	var variantId int64
//	has_results := false
//	variantIdList := []int64{}
//	for results.Next() {
//		has_results = true
//		err = results.Scan(&variantId)
//		if err != nil {
//			return nil, err
//		}
//		variantIdList = append(variantIdList, variantId)
//	}
//	if (!has_results) {
//		return nil, sql.ErrNoRows
//	}
//	return variantIdList, nil
//}
//
//func GetVariantDetail(variantDetailsChan chan VariantDetails, variantId int64,db *sql.DB){
//	variant := VariantDetails{VariantId: variantId}
//	if err := variant.GetVariant(db); err != nil {
//		log.Println(err)
//	} else {
//		variantDetailsChan <- variant
//	}
//
//}
//
//func GetChildVariantsDetailsList(variantIdsList []int64, db *sql.DB) []VariantDetails {
//	variantDetailsChan := make(chan VariantDetails)
//	defer close(variantDetailsChan)
//
//	variantDetailsList := []VariantDetails{}
//
//	for _, variantId := range variantIdsList {
//		go GetVariantDetail(variantDetailsChan,variantId, db)
//	}
//
//	totalVariantChild := len(variantIdsList)
//	for totalVariantChild > 0 {
//		variantDetailsList = append(variantDetailsList, <-variantDetailsChan)
//		totalVariantChild -=1
//	}
//	return variantDetailsList
//}
//
//func DeleteChildVariant(variantDeleteChan chan bool, variantId int64,db *sql.DB) {
//	variant := VariantDetails{VariantId: variantId}
//	if err := variant.DeleteVariant(db); err != nil {
//		log.Println("Error deleting child variant: ",err)
//		variantDeleteChan <- false
//	} else {
//		variantDeleteChan <- true
//	}
//
//}