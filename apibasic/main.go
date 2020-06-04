package main

import (
	"apisingolang/apibasic/operation"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("helloworld Max: received a request")

	//operation.Suma(2, 2)
	fmt.Fprintf(w, "Hello Massimiliano Marocchi: la suma de %d+%d=%d!\n", 3, 3, operation.Suma(3, 3))

}

func main() {
	log.Print("helloworld: starting server...")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
