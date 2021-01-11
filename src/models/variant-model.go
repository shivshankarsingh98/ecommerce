package models

import (
	"database/sql"
	"errors"
	"strconv"
)

type VariantDetails struct {
	VariantId     int64   `json:"variant_id"`
	VariantName   string  `json:"name,omitempty"`
	Mrp           float64   `json:"mrp"`
	DiscountPrice float64   `json:"discount_price,omitempty"`
	Size          string   `json:"size,omitempty"`
	Colour        string  `json:"colour,omitempty"`
	ProductId     int64   `json:"product_id,omitempty"`
}

func (vd *VariantDetails) GetVariant(db *sql.DB) error {

	results, err := db.Query("select * from variant where variant_id = ?", vd.VariantId)
	if err != nil {
		return err
	}
	defer results.Close()

	var pid int
	has_results := false
	for results.Next() {
		has_results = true
		err = results.Scan(&vd.VariantId, &vd.VariantName,& vd.Mrp, &vd.DiscountPrice, &vd.Size, &vd.Colour, &pid)
		if err != nil {
			return err
		}
	}
	if (!has_results) {
		return sql.ErrNoRows
	}
	return nil
}

func (vd *VariantDetails) CreateVariant(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO variant (variant_id, variant_name, mrp,discount_price, size, color, product_id) VALUES (?, ?, ?, ?, ?, ?, ? )")
	if err != nil{
		return err
	}
	_, err = stmt.Exec(vd.VariantId, vd.VariantName, vd.Mrp, vd.DiscountPrice, vd.Size, vd.Colour, vd.ProductId)
	if err != nil {
		return err
	}
	return nil
}

func (vd *VariantDetails) DeleteVariant(db *sql.DB) error {
	stmt, err := db.Prepare("delete from variant where variant_id=?")
	if err != nil{
		return err
	}
	res, err := stmt.Exec(vd.VariantId)
	if err != nil {
		return err
	}else {
		rowEffected, err1 := res.RowsAffected()
		if err1 != nil {
			return err1
		} else if rowEffected == 0 {
			errMsg := "No variant present with id: " +strconv.FormatInt(vd.VariantId, 10)
			return errors.New(errMsg)
		}
	}
	return nil
}


func (vd *VariantDetails) UpdateVariant(db *sql.DB) error {
	var rowEffected int64
	if vd.VariantName != "" {
		stmt, err := db.Prepare("update  variant set variant_name = ? where variant_id= ?")
		if err != nil{
			return err
		}
		res, err := stmt.Exec(vd.VariantName, vd.VariantId)
		rows, _ := res.RowsAffected()
		rowEffected += rows
		if err != nil{
			return err
		}

	}
	if vd.Mrp != -1 {
		stmt, err := db.Prepare("update  variant set mrp = ? where variant_id= ?")
		if err != nil{
			return err
		}
		res, err := stmt.Exec(vd.Mrp, vd.VariantId)
		rows, _ := res.RowsAffected()
		rowEffected += rows
		if err != nil{
			return err
		}

	}
	if vd.DiscountPrice != -1 {
		stmt, err := db.Prepare("update  variant set discount_price = ? where variant_id= ?")
		if err != nil{
			return err
		}
		res, err := stmt.Exec(vd.DiscountPrice, vd.VariantId)
		rows, _ := res.RowsAffected()
		rowEffected += rows
		if err != nil{
			return err
		}

	}
	if vd.Size != "" {
		stmt, err := db.Prepare("update  variant set size = ? where variant_id= ?")
		if err != nil{
			return err
		}
		res, err := stmt.Exec(vd.Size, vd.VariantId)
		rows, _ := res.RowsAffected()
		rowEffected += rows
		if err != nil{
			return err
		}

	}
	if vd.Colour != "" {

		stmt, err := db.Prepare("update  variant set color = ? where variant_id= ?")
		if err != nil{
			return err
		}
		res, err := stmt.Exec(vd.Colour, vd.VariantId)
		rows, _ := res.RowsAffected()
		rowEffected += rows
		if err != nil{
			return err
		}
	}
	if vd.ProductId != -1 {
		stmt, err := db.Prepare("update  variant set product_id = ? where variant_id= ?")
		if err != nil{
			return err
		}
		res, err := stmt.Exec(vd.ProductId, vd.VariantId)
		rows, _ := res.RowsAffected()
		rowEffected += rows
		if err != nil{
			return err
		}

	}

	if rowEffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}