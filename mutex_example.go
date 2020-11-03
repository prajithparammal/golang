package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// Book Struct

type (
	Book struct {
		ID     string  `json:"id"`
		Isbn   string  `json:"isbn"`
		Title  string  `json:"title"`
		Author *Author `json:"author"`
	}

	Author struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	}

	// First we define a type that is a slice of books
	BookStruct struct {
		Books []Book
		m     *sync.Mutex
	}
)

// Next step is - we're gonna tie all our API functions to the new type as if they were methods

func (b *BookStruct) ErrHandler(w http.ResponseWriter, errCode int) {
	w.WriteHeader(errCode)
}

//GET All books
func (b *BookStruct) getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(b.Books); err != nil {
		b.ErrHandler(w, 501)
	}
}

//GET single Book
func (b *BookStruct) getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//loop through books and find ID
	for _, item := range b.Books {
		if item.ID == params["id"] {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				b.ErrHandler(w, 501)
			}
			return
		}
	}
	if err := json.NewEncoder(w).Encode(&Book{}); err != nil {
		b.ErrHandler(w, 501)
	}
}

//Create a new book
func (b *BookStruct) createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	b.m.Lock()
	b.Books = append(b.Books, book)
	if err := json.NewEncoder(w).Encode(b.Books); err != nil {
		b.ErrHandler(w, 501)
	}
	b.m.Unlock()
}

//update book
func (b *BookStruct) updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range b.Books {
		if item.ID == params["id"] {
			b.m.Lock()
			b.Books = append((b.Books)[:index], (b.Books)[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			b.Books = append(b.Books, book)
			if err := json.NewEncoder(w).Encode(b.Books); err != nil {
				b.ErrHandler(w, 501)
			}
			b.m.Unlock()
			return
		}
	}
	if err := json.NewEncoder(w).Encode(b.Books); err != nil {
		b.ErrHandler(w, 501)
	}
}

// Delete a book
func (b *BookStruct) deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range b.Books {
		if item.ID == params["id"] {
			b.m.Lock()
			(b.Books)[index], (b.Books)[len(b.Books)-1] = (b.Books)[len(b.Books)-1], (b.Books)[index]
			b.Books = (b.Books)[:len(b.Books)-1]
			b.m.Unlock()
			break
		}
	}
	if err := json.NewEncoder(w).Encode(b.Books); err != nil {
		b.ErrHandler(w, 501)
	}
}

func main() {
	var x BookStruct
	// Init Router
	r := mux.NewRouter()

	// Init Books vars as a slice Book struct
	x.Books = append(x.Books, Book{ID: "1", Isbn: "1000", Title: "harrypotter", Author: &Author{Firstname: "J.K", Lastname: "Rowling"}})
	x.Books = append(x.Books, Book{ID: "2", Isbn: "2000", Title: "Mahabarth", Author: &Author{Firstname: "Maha", Lastname: "Ganpathi"}})
	x.m = &sync.Mutex{}
	// Router Handler / Endpoints
	r.HandleFunc("/api/books", x.getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", x.getBook).Methods("GET")
	r.HandleFunc("/api/books", x.createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", x.updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", x.deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
