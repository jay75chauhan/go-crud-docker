package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	// "github.com/lib/pq"
)


type User struct {
	Id   int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`

}

func main() {
	//conect the database

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	if err!= nil {
    panic(err)
  }
	defer db.Close()

	//create router
	r := mux.NewRouter()
  r.HandleFunc("/users", getUser(db)).Methods("GET")
  r.HandleFunc("/users/{id}", getUser(db)).Methods("GET")
  r.HandleFunc("/users", createUser(db)).Methods("POST")
  r.HandleFunc("/users/{id}", updateUser(db)).Methods("PUT")
  r.HandleFunc("/users/{id}", deletUser(db)).Methods("DELETE")


	//start server

	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router)))

}

func jsonContentTypeMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    next.ServeHTTP(w, r)
  })
}