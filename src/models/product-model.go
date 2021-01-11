package models

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

type ProductDetails struct {
	ProductId         int64              `json:"product_id"`
	ProductName       string             `json:"name"`
	Description       string             `json:"description,omitempty"`
	ProductImageUrl   string             `json:"product_image_url,omitempty"`
	ChildVariants     []VariantDetails   `json:"variants"`
	CategoryIds       []int64            `json:"category_ids,omitempty"`
}

func (pd *ProductDetails) GetProduct(db *sql.DB) error {
	results, err := db.Query("select * from product where product_id = ?", pd.ProductId)
	if err != nil {
		return err
	}
	defer results.Close()

	has_results := false
	for results.Next() {
		has_results = true
		err = results.Scan(&pd.ProductId, &pd.ProductName, &pd.Description, &pd.ProductImageUrl)
		if err != nil {
			return err
		}
	}
	if (!has_results) {
		return sql.ErrNoRows
	}

	variantIdsList, err := GetChildVariantsIdsList(db,pd.ProductId)
	if err != nil{
		return nil
	}
	ChildVariants := GetChildVariantsDetailsList(variantIdsList, db)
	pd.ChildVariants = ChildVariants
	return nil
}

func (pd *ProductDetails) CreateProduct(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO product (product_id, product_name, description,url) VALUES (?, ?, ?, ?)")
	if err != nil{
		return err
	}
	_, err = stmt.Exec(pd.ProductId, pd.ProductName, pd.Description, pd.ProductImageUrl)
	if err != nil {
		return err
	}

	insertCheck := make(chan bool)
	defer close(insertCheck)

	for _, categortId :=  range pd.CategoryIds {
		go CreateProductCategoryRelation(db,pd.ProductId,categortId,insertCheck)
	}

	totalCategoryId := len(pd.CategoryIds)

	for totalCategoryId > 0 {
		check := <- insertCheck
		if check == false {
			return errors.New("Error while inserting into product_category table")
		}
		totalCategoryId -= 1
	}

	return nil
}

func CreateProductCategoryRelation(db *sql.DB, productId, categoryId int64, insertCheck chan bool) {
	stmt, err := db.Prepare("INSERT INTO product_category (product_id, category_id) VALUES (?, ?)")
	if err != nil{
		log.Println("Error in  product-category relation query: ", err)
		insertCheck <- false
	}
	_, err = stmt.Exec(productId, categoryId)
	if err != nil {
		log.Println("Error while executing product-category relation query: ", err)
		insertCheck <- false
	}
	insertCheck <- true
}

func (pd *ProductDetails) DeleteProduct(db *sql.DB) error {
	variantDeleteChan := make(chan bool)
	defer close(variantDeleteChan)

	variantIdsList, err := GetChildVariantsIdsList(db,pd.ProductId)
	for _, variantId := range variantIdsList {
		go DeleteChildVariant(variantDeleteChan,variantId, db)
	}

	totalVariantChild := len(variantIdsList)
	for totalVariantChild > 0 {
		check := <- variantDeleteChan
		if check == false {
			return errors.New("Error while deleting variant of productId: "+ strconv.FormatInt(pd.ProductId, 10))
		}
		totalVariantChild -= 1
	}

	stmt, err := db.Prepare("delete from product where product_id=?")
	if err != nil{
		return err
	}
	res, err := stmt.Exec(pd.ProductId)
	if err != nil {
		return err
	}else {
		rowEffected, err1 := res.RowsAffected()
		if err1 != nil {
			return err1
		} else if rowEffected == 0 {
			errMsg := "No product present with id: " +strconv.FormatInt(pd.ProductId, 10)
			return errors.New(errMsg)
		}
	}

	deleteFromProductCategorystmt, err := db.Prepare("delete from product_category where product_id=?")
	if err != nil{
		return err
	}
	res1, err := deleteFromProductCategorystmt.Exec(pd.ProductId)
	if err != nil {
		return err
	}else {
		rowEffected, err1 := res1.RowsAffected()
		if err1 != nil {
			return err1
		} else if rowEffected == 0 {
			errMsg := "No product present with id: " +strconv.FormatInt(pd.ProductId, 10) + " in product_category table"
			return errors.New(errMsg)
		}
	}
	return nil
}

func (pd *ProductDetails) UpdateProduct(db *sql.DB) error {
	if pd.ProductName != "" {
		stmt, err := db.Prepare("update  product set product_name = ? where product_id= ?")
		if err != nil{
			return err
		}
		_, err = stmt.Exec(pd.ProductName, pd.ProductId)
		if err != nil{
			return err
		}
	}
	if pd.Description != "" {
		stmt, err := db.Prepare("update  product set description = ? where product_id= ?")
		if err != nil{
			return err
		}
		_, err = stmt.Exec(pd.Description, pd.ProductId)
		if err != nil{
			return err
		}
	}
	if pd.ProductImageUrl != "" {
		stmt, err := db.Prepare("update  product set url = ? where product_id= ?")
		if err != nil{
			return err
		}
		_, err = stmt.Exec(pd.ProductImageUrl, pd.ProductId)
		if err != nil{
			return err
		}
	}
	return nil
}

func GetChildVariantsIdsList(db *sql.DB, productId int64) ([]int64,error) {
	results, err := db.Query("select variant_id from variant where product_id = ?", productId)
	if err != nil {
		return nil, err
	}
	defer results.Close()

	var variantId int64
	has_results := false
	variantIdList := []int64{}
	for results.Next() {
		has_results = true
		err = results.Scan(&variantId)
		if err != nil {
			return nil, err
		}
		variantIdList = append(variantIdList, variantId)
	}
	if (!has_results) {
		return nil, sql.ErrNoRows
	}
	return variantIdList, nil
}

func GetVariantDetail(variantDetailsChan chan VariantDetails, variantId int64,db *sql.DB){
	variant := VariantDetails{VariantId: variantId}
	if err := variant.GetVariant(db); err != nil {
		log.Println(err)
	} else {
		variantDetailsChan <- variant
	}

}

func GetChildVariantsDetailsList(variantIdsList []int64, db *sql.DB) []VariantDetails {
	variantDetailsChan := make(chan VariantDetails)
	defer close(variantDetailsChan)

	variantDetailsList := []VariantDetails{}

	for _, variantId := range variantIdsList {
		go GetVariantDetail(variantDetailsChan,variantId, db)
	}

	totalVariantChild := len(variantIdsList)
	for totalVariantChild > 0 {
		variantDetailsList = append(variantDetailsList, <-variantDetailsChan)
		totalVariantChild -=1
	}
	return variantDetailsList
}

func DeleteChildVariant(variantDeleteChan chan bool, variantId int64,db *sql.DB) {
	variant := VariantDetails{VariantId: variantId}
	if err := variant.DeleteVariant(db); err != nil {
		log.Println("Error deleting child variant: ",err)
		variantDeleteChan <- false
	} else {
		variantDeleteChan <- true
	}

}