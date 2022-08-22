package main

import (
	"fmt"
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

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/index.html")
}

func main() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	http.HandleFunc("/random", handler)
	http.HandleFunc("/", index)
	fmt.Println("Hosting a local server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}