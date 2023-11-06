package main

import(
	"fmt"
	"sync"
	"time"
	"net/http"
)

func generate_load(n int) *sync.WaitGroup{
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
	}

	return &wg
}

func test(wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Post("http://localhost:8080/process", "application/json", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	fmt.Println("Starting load test")
	n := 1000
	wg := generate_load(n)

	t1 := time.Now()

	for i := 0; i < n; i++ {
		go test(wg)
	}

	wg.Wait()
	fmt.Println("Time taken :", time.Now().Sub(t1))
}