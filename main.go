package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"firstName"`
	Lastname  string `json:"Lastname"`
}

var movies = []Movie{
	{
		ID:    "1",
		Isbn:  "1001",
		Title: "Inception",
		Director: &Director{
			FirstName: "Christopher",
			Lastname:  "Nolan",
		},
	},
	{
		ID:    "2",
		Isbn:  "1002",
		Title: "Interstellar",
		Director: &Director{
			FirstName: "Christopher",
			Lastname:  "Nolan",
		},
	},
	{
		ID:    "3",
		Isbn:  "1003",
		Title: "The Dark Knight",
		Director: &Director{
			FirstName: "Christopher",
			Lastname:  "Nolan",
		},
	},
	{
		ID:    "4",
		Isbn:  "1004",
		Title: "Titanic",
		Director: &Director{
			FirstName: "James",
			Lastname:  "Cameron",
		},
	},
	{
		ID:    "5",
		Isbn:  "1005",
		Title: "Avatar",
		Director: &Director{
			FirstName: "James",
			Lastname:  "Cameron",
		},
	},
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
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
	json.NewEncoder(w).Encode(movies)
}

func getMoviebyid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}

}
func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func updatemovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movies)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return

		}

	}

}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMoviebyid).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")

	fmt.Printf("server start at PORT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
