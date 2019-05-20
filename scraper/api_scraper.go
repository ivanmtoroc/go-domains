package scraper

import (
	"log"
	"time"
	"net/http"
	"encoding/json"
	"go-domains/models"
	"go-domains/utilities"
)

const API_URL = "https://api.ssllabs.com/api/v3/analyze?host="
var GRADES = map[string]int {
	"A+":	1, "A":	2, "B":	3, "C":	4, "D":	5, "E":	6, "F":	7, "Invalid": 8,
}

func GetServers(domain *models.Domain) []*models.Server {
	var servers []*models.Server

  result := consultAPI(domain)
  if domain == nil {
    return nil
  } else if domain.IsDown {
		return servers
  }

	servers = getServerToAPI(domain, result)
	return servers
}

func consultAPI(domain *models.Domain) map[string]interface{} {
	var result map[string]interface{}

  log.Println("Consulting API - Host: " + domain.Name)
	for {
		response, err := http.Get(API_URL + domain.Name)
		if err != nil {
      domain = nil
      return nil
		}

		err = json.NewDecoder(response.Body).Decode(&result)
		if err == nil {
			status := result["status"].(string)
			log.Println("  - Status: " + status)
			if status == "READY" {
				return result
			} else if status == "ERROR" {
				domain.IsDown = true
	      return nil
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func getServerToAPI(domain *models.Domain, result map[string]interface{}) []*models.Server {
	var servers []*models.Server

	domain.SslGrade = "A+"

	for _, element := range result["endpoints"].([]interface{}) {
		endpoint := element.(map[string]interface{})

		ip := endpoint["ipAddress"].(string)

		grade := "Invalid"
		status := endpoint["statusMessage"].(string)
		if status == "Ready" {
			grade = endpoint["grade"].(string)
		}
		if GRADES[grade] > GRADES[domain.SslGrade] {
			domain.SslGrade = grade
		}

		command := "whois " + ip + " | grep -E '(C|c)ountry' | head -n 1 | sed -E 's/(C|c)ountry: *//'"
		country := utilities.RunShCommant(command)

		command = "whois " + ip + " | grep -E 'OrgName|org-name|owner' | head -n 1 | sed -E 's/(OrgName|org-name|owner): *//'"
		owner := utilities.RunShCommant(command)

		server := &models.Server{ip, grade, country, owner, domain.Name, domain.CreatedAt}
		servers = append(servers, server)
	}

	return servers
}
