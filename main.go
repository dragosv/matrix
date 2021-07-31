package main

import (
	"github.com/dragosv/matrix/handlers"
	"github.com/swaggo/http-swagger"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs" // docs are generated by Swag CLI, have to import it.
	"net/http"
)

// @title Matrix web server API
// @version 1.0
// @description Basic web server written in GoLang that performs operations on a matrix.
// @contact.name Dragos Varovici
// @contact.url https://github.com/dragosv
// @license.name Mozilla Public License 2.0
// @license.url https://www.mozilla.org/en-US/MPL/2.0/
// @host localhost:8080
func main() {
	http.HandleFunc("/echo", handlers.EchoHandler)
	http.HandleFunc("/invert", handlers.InvertHandler)
	http.HandleFunc("/flatten", handlers.FlattenHandler)
	http.HandleFunc("/sum", handlers.SumHandler)
	http.HandleFunc("/multiply", handlers.MultiplyHandler)
	http.HandleFunc("/swagger", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	http.ListenAndServe(":8080", nil)
}
