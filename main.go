package main

import (
  "net/http"
  "go-domains/handlers"
  "go-domains/models"
  "github.com/go-chi/chi"
  "github.com/go-chi/render"
)

func main() {
  models.InitDB()

	router := chi.NewRouter()

	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Wellcome!"))
	})

	router.Route("/domains", func(router chi.Router) {
		router.Get("/{domain_name}", handlers.GetDomain)
	})

	http.ListenAndServe(":3333", router)
}
