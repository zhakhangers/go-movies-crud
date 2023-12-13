package main

import (
	"fmt"
	"log"
	"strconv"
	"github.com/gorilla/mux"
	"encoding/json"
	"math/rand"
	"net/http"
)

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies{

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies{

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)

}


func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1010000))
	movies = append(movies,movie)
	json.NewEncoder(w).Encode(movie)

}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies{

		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(1010000))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}

	}

}






func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID:"1", Isbn:"912836", Title:"Wonderboy", Director: &Director{Name:"Steven", Surname:"Thompson"}})
	movies = append(movies, Movie{ID:"2", Isbn:"802509", Title:"Afghani", Director: &Director{Name:"Belal", Surname:"Mukhammed"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))


}