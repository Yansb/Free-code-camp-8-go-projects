package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yansb/go-bookstore/pkg/routes"
)

func hello(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(map[string]string{"message": "Hello World"})
	response.Write(res)
}

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	fmt.Println("Server is running on port 9010")
	r.HandleFunc("/", hello)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":9010", r))
}
