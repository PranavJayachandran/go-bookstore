package main

import (
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/PranavJayachandran/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r:=mux.NEwRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:8000",r))
	
}