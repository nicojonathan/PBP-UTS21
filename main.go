package main

import (
	"fmt"      // used for formatted printing
	"log"      // used for logging
	"net/http" //  standard library package that provides functionality for building HTTP servers and clients.
	"uts21/controllers"

	_ "github.com/go-sql-driver/mysql" //  is a blank import to ensure that the MySQL driver is included in the build.
	"github.com/gorilla/mux"           // is used for routing and handling HTTP requests.
)


func main() {
	router := mux.NewRouter()

	// endpoint no1
	router.HandleFunc("/user/login", controllers.Login).Methods("GET")

	//endpoint no2
	router.HandleFunc("/song/popular", controllers.GetPopularSongs).Methods("GET")

	//endpoint no 3
	router.HandleFunc("/song/recommended", controllers.GetRecommendedSong).Methods("GET")
	

	http.Handle("/", router)

	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}