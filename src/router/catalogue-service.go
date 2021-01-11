package router

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"

)

type ProductCatalogueService struct {
	Router *mux.Router
	DB     *sql.DB
}

func (pcs *ProductCatalogueService) Initialize(user, password, dbname string) {
	var err error
	dbDriver := "mysql"
	dbUser := user
	dbPass := password
	dbName := dbname

	pcs.DB, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal("Error opening mysql client: ", err)
		return
	}else {
		if err := pcs.DB.Ping(); err!= nil {
			log.Println("Mysql ping err: ",err)
			return
		}else {
			log.Println("Service sucessefully connected to Mysql client ")
		}
	}

	pcs.Router = mux.NewRouter()
	pcs.initializeCategoryRoutes()
	pcs.initializeProductRoutes()
	pcs.initializeVariantRoutes()

}


func (pcs *ProductCatalogueService) RunService(addr string) {
	log.Println("Product catalogue service started")
	log.Fatal(http.ListenAndServe(addr, pcs.Router))
}