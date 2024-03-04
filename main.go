package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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

//	http.ResponseWriter ->
// sending HTTP responses to a client. It allows you to write your response data and headers.
//r *http.Request in Go, think of it like this:
//You have a note (r) that is actually a map (*http.Request)
//to find a very special toy box (http.Request).
//This toy box has all sorts of things you can ask for when you're playing on the internet, like asking for a new webpage to look at or sending a letter to a website.

func getMovies(w http.ResponseWriter, r *http.Request) {
	//Content-Type header of the response to "application/json".
	//This tells the client that the server is returning JSON-formatted data.
	w.Header().Set("Content Type", "application/json")
	//that will write output to w, the http.ResponseWriter.
	//This encoder can convert Go data structures into JSON format

	//Takes a Go data structure (movies in this case, which is likely a slice or array of
	//movie-related structs) and encodes it as JSON.
	//This encoded JSON data is then written directly to the http.ResponseWriter,
	//effectively sending it as the HTTP response body.
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

// so remember the w here is the tool to send things back to the user
// r here is the actual information received by the request

func getMovie(w http.ResponseWriter, r *http.Request) {
	// The response will be in json
	w.Header().Set("Content-Type", "application/json")
	// mux is a package often used in Go for dealing with HTTP requests
	// Vars is a function from the mux package. Its job is to extract variables from the HTTP request.
	params := mux.Vars(r)
	// _ mean it does not need this value and we're going to use the item which is the value of
	//the current element and not the index
	// Because remember in GO it give an error when not using a variable
	for _, item := range movies {
		if item.ID == params["id"] {
			// If the movie is found (its ID matches), this line sends back the
			//information about the movie in JSON format.
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
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
