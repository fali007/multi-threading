package main

import (
	"fmt"
	"log"
	"net/http"
	"math/rand"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Person struct {
	Name string
	Age int
	Address string
	Job string
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func to_json(i interface{}) []byte {
	json,err:=json.Marshal(i)
	if err!=nil{
		fmt.Println("Error marshalling json :",err)
	}
	return json
}

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func generate_sample () Person {
	name := randStringBytes(12)
	age := rand.Intn(70)
	address := randStringBytes(20)
	job := randStringBytes(8)
	return Person{name, age, address, job}
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(to_json(generate_sample()))
}

func main() {
	fmt.Println("Starting server")

	r := mux.NewRouter()
	r.HandleFunc("/process", processRequest).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}