package controllers

import (
	"time"
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

		w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if domainName := chi.URLParam(r, "domainName"); domainName != "" {
			domain, servers = scraper.GetDomainAPI(domainName)
      if servers == nil {
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

	previous_domain := models.GetDomainDB(domain.Name)

	domain.Create()
	for _, server := range servers {
		server.Create()
	}

	one_hour, _ := time.ParseDuration("1h")
	previous_domain_time := previous_domain.UpdatedAt.Add(one_hour)
	domain_time := domain.UpdatedAt

	if previous_domain_time.Before(domain_time) || previous_domain_time.Equal(domain_time) {
		domain.ServersChanged = true
		domain.PreviousSslGrade = previous_domain.SslGrade
	}

	if err := render.Render(w, r, api.CreateDomainResponse(domain, servers)); err != nil {
		render.Render(w, r, api.ERR404)
		return
	}
}
