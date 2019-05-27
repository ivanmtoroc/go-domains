package main

import (
  "log"
  "net/http"
  "go-domains/handlers"
  "go-domains/models"
  "github.com/go-chi/chi"
  "github.com/go-chi/render"
)

func main() {
  // Start connection with data base
  models.InitDB()
  models.CreateTables()

  router := chi.NewRouter()

  // Set ContentType: 'aplication/json' in header of responses
  router.Use(render.SetContentType(render.ContentTypeJSON))

  // Set routes
  router.Route("/api/v1/", func(router chi.Router) {
  router.Get("/domains/{domain_name}", handlers.GetDomain)
    router.Get("/items", handlers.GetItems)
  })

  log.Println("Server listening...")
  http.ListenAndServe(":3333", router)
}
