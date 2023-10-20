package main

import (
	"fmt"
	"net/http"
)

func runServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "First WebService in Golang")
	})

	var port int
	port = 9003
	fmt.Printf("Server is running on port %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		fmt.Println("Error starting WebService: ", err)
	}
}

func main() {
	runServer()
}