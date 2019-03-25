package main

import (
	"fmt"
	"github.com/ziellim/cuvi/model"
	"net/http"
)

func main() {
	http.HandleFunc("/", getOne)
	http.ListenAndServe(":8080", nil)
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
