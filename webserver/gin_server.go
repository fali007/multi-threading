package main

import (
	"fmt"
	"time"
	"net/http"
	"math/rand"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string
	Age int
	Address string
	Job string
}

var req_channel chan *gin.Context

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

func processRequest(c *gin.Context) {
	req_channel <- c
}

func thread_func(c *gin.Context) {
	c.Data(http.StatusOK, gin.MIMEJSON, to_json(generate_sample()))
}

func process_thread(c *chan *gin.Context) {
	counter := 0
	for i := range *c {
		go thread_func(i)
		counter ++
		if counter % 10 == 0 {
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	fmt.Println("Starting server")
	req_channel = make(chan *gin.Context)

	go process_thread(&req_channel)

	w := gin.Default()
	w.POST("/process", processRequest)
	w.Run()
}