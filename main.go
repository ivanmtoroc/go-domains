package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/ivanmtoroc/go-domains/handlers"
	"github.com/ivanmtoroc/go-domains/models"
)

func main() {
	// Start connection with database
	models.InitDB()
	// Create tables into database
	models.CreateTables()

	// Create new router
	router := chi.NewRouter()

	// Add middledware to set "ContentType: 'aplication/json'" in all responses
	router.Use(render.SetContentType(render.ContentTypeJSON))

	// Define routes and add handlers functions
	router.Route("/api/v1/", func(router chi.Router) {
		// "/domains/{domainName}" get all info to domain
		router.Get("/domains/{domainName}", handlers.GetDomain)
		// "/items" get domains history
		router.Get("/items", handlers.GetItems)
	})

	// Start server in "localhost:3333"
	fmt.Println("Server listening...")
	http.ListenAndServe(":3333", router)
}
