package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./images/")
	if err != nil {
		log.Fatal(err)
	}
	file := files[rand.Intn(len(files))].Name()

	http.ServeFile(w, r, "./images/"+file)
}

func main() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}