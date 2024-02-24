package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// like a box where you can store different kinds of things together,
// like a person's name and age. It helps you keep related information organized in one place
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` // The * says it might be null
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	if err != nil {
		return
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break

		}
	}

}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{
		ID:       "1",
		Isbn:     "438227",
		Title:    "Movie one",
		Director: &Director{Firstname: "John", Lastname: "Doe"}})

	movies = append(movies, Movie{
		ID:       "2",
		Isbn:     "438455",
		Title:    "Movie two",
		Director: &Director{Firstname: "Steve", Lastname: "smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000 \n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
