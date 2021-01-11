package main

import (
	"router"
)

var (
	mysqlUser = "root"
	mysqlPass = "Pass1234"
	databaseName = "ecommerce"
)

func main(){
	productCatalogue := router.ProductCatalogueService{}
	productCatalogue.Initialize(mysqlUser,mysqlPass,databaseName)
	productCatalogue.RunService(":8080")
}


