package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/felix1251/go-rest-api/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterPostRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}

