package controllers

import (
	"context"
	"net/http"
	"go-domains/api"
	"go-domains/models"
	"go-domains/scraper"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func DomainCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var domain *models.Domain
		var servers []*models.Server
		var err error

		if domainName := chi.URLParam(r, "domainName"); domainName != "" {
			domain, servers = scraper.GetDomainAPI(domainName)
      if err != nil || domain == nil {
  			render.Render(w, r, api.ERR404)
  			return
  		}
		} else {
			render.Render(w, r, api.ERR404)
			return
		}

		my_context := context.WithValue(r.Context(), "domain", domain)
		my_context = context.WithValue(my_context, "servers", servers)
		next.ServeHTTP(w, r.WithContext(my_context))
	})
}

func GetDomain(w http.ResponseWriter, r *http.Request) {
	domain := r.Context().Value("domain").(*models.Domain)
	servers := r.Context().Value("servers").([]*models.Server)
	if err := render.Render(w, r, api.CreateDomainResponse(domain, servers)); err != nil {
		render.Render(w, r, api.ERR404)
		return
	}
}
