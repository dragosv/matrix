package main

import (
	"github.com/dragosv/matrix/handlers"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	http.HandleFunc("/echo", handlers.EchoHandler)
	http.HandleFunc("/invert", handlers.InvertHandler)
	http.HandleFunc("/flatten", handlers.FlattenHandler)
	http.HandleFunc("/sum", handlers.SumHandler)
	http.HandleFunc("/multiply", handlers.MultiplyHandler)

	http.ListenAndServe(":8080", nil)
}
