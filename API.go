package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete ","status":"100"}`))
}

func updateMovies(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "update response ","status":"100"}`))
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Create movies",  "status":"0"}`))
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"get movies by ID"}, "status":"0"`))
}

func main() {

	fmt.Println("Testing the API")

	r := mux.NewRouter()

	movies = append(movies, Movies{"893737", "Isbn673", "Rich Dad Poor Dad",
		"Financial Literacy ",
		&Director{"6553", "Michael Loyloey", "gogthe@yahoo.com", "020827625"}})

	movies = append(movies, Movies{"873590", "Isbn7836", "Who move my Cheese",
		"Business Strategy",
		&Director{"00176", "Isaac Recon", "ishjeu@gmail.com", "020872652"}})

	movies = append(movies, Movies{"9876", "Isbn9876", "Who move my Cheese Revenge",
		"Business Strategy Blue sea",
		&Director{"00176", "Isaac Recon", "ishjeu@gmail.com", "020872652"}})

	movies = append(movies, Movies{"28726", "Isbn6897", "The Lord of the Wings",
		"Strategic Decision Making ",
		&Director{"00176", "Isaac Recon", "ishjeu@gmail.com", "020872652"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.Fatal(err)
	}
}
