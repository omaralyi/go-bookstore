package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/omaralyi/go-bookstore/pkg/routes"
)
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	r.
}
