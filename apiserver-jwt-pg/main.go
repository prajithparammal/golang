/*
Creating users in pq:

create table users (
   id serial primary key,
   email text not null unique,
   password text not null
);

insert into users(email,password) values ('test@example.com', '12345')

*/

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"` // need to decode user provided payload to User struct
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"` // we need to pass it to user as json if user is authenticated
}

type Error struct {
	Message string `json:"message"` // Need to pass it to end user as json when there is any error
}

var db *sql.DB

func main() {
	pgUrl, err := pq.ParseURL("postgres://wiykjqny:8LTkzjTczQ5LVpGkbhxcl5saZ65Kwzla@lallah.db.elephantsql.com:5432/wiykjqny")
	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	fmt.Println(err)

	router := mux.NewRouter()
	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleware(protectedEndpoint)).Methods("GET")

	log.Println("Listening on Port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))

}

/////////////////////////////////////////
func respondWithError(w http.ResponseWriter, status int, error Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func reponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

///////////////////////////////////////////
func signup(w http.ResponseWriter, r *http.Request) {
	var user User
	var error Error
	_ = json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email is missing"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "Password is missing"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(hash)

	stmt := "insert into users (email, password) values($1, $2) RETURNING id;"
	err = db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		error.Message = "Server error"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	reponseJSON(w, user)
}

//////////////////////////////////////////////////////

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login invoked")

}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signup invoked")

}

func TokenVerifyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return nil
}
