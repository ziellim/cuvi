package main

import (
	"fmt"
	"github.com/ziellim/cuvi/model"
	"net/http"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	})
	router := mux.NewRouter()
	router.HandleFunc("/v0/cvs/0", getOne).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(router)))
	http.ListenAndServe(":8080", c.Handler(router))

}

func getOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(getCVAsJson())
}

func getCVAsJson() []byte {
	cv := model.LoadCV()
	json, err := cv.Marshal()
	if err != nil {
		fmt.Println(err)
	}
	return json
}
