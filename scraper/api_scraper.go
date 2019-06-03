package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ivanmtoroc/go-domains/models"
	"github.com/ivanmtoroc/go-domains/utilities"
)

// APIURL is the SSL Labs API url
const APIURL = "https://api.ssllabs.com/api/v3/analyze?host="

// Grades options by SSL Labs
var Grades = map[string]int{
	"A+":      1,
	"A-":      2,
	"A":       3,
	"B":       4,
	"C":       5,
	"D":       6,
	"E":       7,
	"F":       8,
	"T":       9,
	"M":       10,
	"Unknown": 11, // To down domains
}

// getServers get all servers of a domain using SSL Labs API
func getServers(domain *models.Domain) ([]*models.Server, error) {
	var servers []*models.Server
	// Get JSON response to API
	jsonResult, err := consultAPI(domain)
	if err != nil {
		return nil, err
	} else if !domain.IsValid {
		return nil, nil
	}
	// Get servers from JSON response and set SslGrade to domain
	servers = getServerFromResponse(domain, jsonResult)
	return servers, nil
}

// consultAPI fetch JSON response from SSL Labs API by domain
func consultAPI(domain *models.Domain) (map[string]interface{}, error) {
	var jsonResult map[string]interface{}

	fmt.Println("consulting SSL Labs API")
	for {
		// Make http request to API
		response, err := http.Get(APIURL + domain.Name)
		if err != nil {
			log.Println("API request error")
			log.Println("- error: ", err)
			return nil, err
		}

		// Decode response
		err = json.NewDecoder(response.Body).Decode(&jsonResult)
		if err == nil {
			status := jsonResult["status"].(string)
			fmt.Println("- status: " + status)
			if status == "READY" {
				return jsonResult, nil
			} else if status == "ERROR" {
				domain.IsValid = false
				return nil, nil
			}
		}
		// Wait three seconds to next request
		time.Sleep(3 * time.Second)
	}
}

// getServerFromResponse extract endpoints (servers) to JSON result
func getServerFromResponse(domain *models.Domain, jsonResult map[string]interface{}) []*models.Server {
	var servers []*models.Server
	var command string
	var err error

	// By default domain have the best SSL grade and is down
	domain.SslGrade = "A+"
	domain.IsDown = true

	fmt.Println("get endpoints to domain")
	// Iterate endpoints of JSON result
	for i, item := range jsonResult["endpoints"].([]interface{}) {
		// Type assertion to map
		endpoint := item.(map[string]interface{})
		fmt.Printf("- endpoint %d\n", i)

		// Create new server
		server := &models.Server{
			SslGrade: "Unknown",
		}

		// Get IP address of endpoint
		server.Address = endpoint["ipAddress"].(string)
		fmt.Printf(" - ip: %s\n", server.Address)

		// By default all enpoints have invalid SSL (is down)
		status := endpoint["statusMessage"].(string)
		// If status is "Ready" get SSL grade
		if status == "Ready" {
			server.SslGrade = endpoint["grade"].(string)
			// If at least one enpoint is up then domain is not down
			domain.IsDown = false
		}
		fmt.Printf(" - grade: %s\n", server.SslGrade)

		// Update min SSL Grade of domain
		if Grades[server.SslGrade] > Grades[domain.SslGrade] {
			domain.SslGrade = server.SslGrade
		}

		// Execute whois command with IP of server to get country
		command = "whois " + server.Address + " | grep -E '(C|c)ountry' | head -n 1 | sed -E 's/(C|c)ountry: *//'"
		if server.Country, err = utilities.RunShCommant(command); err != nil {
			server.Country = "Country not found"
		}
		fmt.Printf(" - country: %s\n", server.Country)

		// Execute whois command with IP of server to get owner
		command = "whois " + server.Address + " | grep -E 'OrgName|org-name|owner' | head -n 1 | sed -E 's/(OrgName|org-name|owner): *//'"
		if server.Owner, err = utilities.RunShCommant(command); err != nil {
			server.Owner = "Owner not found"
		}
		fmt.Printf(" - owner: %s\n", server.Owner)

		// Add server to servers collection
		servers = append(servers, server)
	}

	return servers
}
