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

type Movies struct {
	Id          string    `json:"id"`
	Isbn        string    `json:"isbn"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Director    *Director `json:"director"`
}

type Director struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
}

var movies []Movies

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, Item := range movies {
		if Item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

func updateMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, Item := range movies {

		if Item.Id != params["id"] {
			w.Write([]byte(`{"message":"Item not found.Provide the right index"}`))
			break
		}

		movies = append(movies[index:], movies[index+1:]...)
		var movie Movies
		_ = json.NewDecoder(r.Body).Decode(&movie)
		movie.Id = params["id"]
		movies = append(movies, movie)
		json.NewEncoder(w).Encode(movies)
		return
	}
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.Id = strconv.Itoa(rand.Intn(1000000000000))
	movie.Director.Id = strconv.Itoa(rand.Intn(289999999999))
	movies = append(movies, movie)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, Item := range movies {
		if Item.Id == params["id"] {
			json.NewEncoder(w).Encode(Item)
			return
		}
	}
}

func main() {

	fmt.Println("Testing the API")

	r := mux.NewRouter()

	movies = append(movies, Movies{"893737", "Isbn673", "Rich Dad Poor Dad",
		"Financial Literacy ",
		&Director{"6553", "Michael Lorey", "gogthe@yahoo.com", "020827625"}})

	movies = append(movies, Movies{"873590", "Isbn7836", "Who move my Cheese",
		"Business Strategy",
		&Director{"00176", "Isaac Recon", "ishjeu@gmail.com", "020872652"}})

	movies = append(movies, Movies{"9876", "Isbn9876", "Who move my Cheese Revenge",
		"Business Strategy Blue sea",
		&Director{"00176", "Isaac Recon", "ishjeu@gmail.com", "020872652"}})

	movies = append(movies, Movies{"28726", "Isbn6897", "The Lord of the Wings",
		"Strategic Decision Making ",
		&Director{"00176", "Isaac Recon", "ishjeu@gmail.com", "020872652"}})

	r.HandleFunc("/api/v1/movies", getMovies).Methods("GET")
	r.HandleFunc("/api/v1/movie/{id}", getMovie).Methods("GET")
	r.HandleFunc("/api/v1/movies", createMovies).Methods("POST")
	r.HandleFunc("/api/v1/movie/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/api/v1/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Go server is Running of port :8080")

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}
}
