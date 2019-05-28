package scraper

import (
  "fmt"
  "log"
  "time"
  "net/http"
  "encoding/json"
  "go-domains/models"
  "go-domains/utilities"
)

// SSL Labs API URL
const API_URL = "https://api.ssllabs.com/api/v3/analyze?host="
// Grades options in SSL Labs
var GRADES = map[string]int {
  "A+": 1, "A-": 2, "A": 3, "B": 4, "C": 5, "D": 6,
  "E": 7, "F": 8, "T": 9, "M": 10, "Unknown": 11,
}

// Function to get servers of domain by SSL Labs API
func getServers(domain *models.Domain) ([]*models.Server, error) {
  var servers []*models.Server

  // Get JSON response to API
  result, err := consultAPI(domain)
  // Validate if domain is not valid or is down
  if err != nil {
    return nil, err
  } else if !domain.IsValid {
    return servers, nil
  }

  // Get servers from JSON response and set SslGrade to domain
  servers = getServerFromResponse(domain, result)

  return servers, nil
}

// Function that fetch JSON response from SSL Labs API by domain
func consultAPI(domain *models.Domain) (map[string]interface{}, error) {
  var result map[string]interface{}

  fmt.Println("consulting API")
  for {
    // Make http request (Method Get) to API
    response, err := http.Get(API_URL + domain.Name)
    if err != nil {
      log.Println("API request error")
      log.Println("- error: ", err)
      return nil, err
    }

    // Decode response
    err = json.NewDecoder(response.Body).Decode(&result)
    if err == nil {
      // Validate status of domain
      status := result["status"].(string)
      fmt.Println("- status: " + status)
      if status == "READY" {
        return result, nil
      } else if status == "ERROR" {
        domain.IsValid = false
        return nil, nil
      }
    }
    // Wait three seconds to next request
    time.Sleep(3 * time.Second)
  }
}

// Function to extract endpoints to API result
func getServerFromResponse(domain *models.Domain, result map[string]interface{}) []*models.Server {
  var servers []*models.Server

  // By default domain have the best SSL grade and is down
  domain.SslGrade = "A+"
  domain.IsDown = true

  fmt.Println("get endpoints to domain")

  // Iterate endpoints of JSON response
  for i, element := range result["endpoints"].([]interface{}) {
    // Type assertion to map
    endpoint := element.(map[string]interface{})
    fmt.Printf("- endpoint %d\n", i)

    // Get IP of server
    ip := endpoint["ipAddress"].(string)
    fmt.Printf(" - ip: %s\n", ip)

    // By default all servers have invalid SSL
    grade := "Unknown"
    status := endpoint["statusMessage"].(string)
    // If status is Ready get SSL grade
    if status == "Ready" {
      grade = endpoint["grade"].(string)
      // If at least one server is up then domain is not down
      domain.IsDown = false
    }
    fmt.Printf(" - grade: %s\n", grade)
    // Update min SSL Grade of domain
    if GRADES[grade] > GRADES[domain.SslGrade] {
      domain.SslGrade = grade
    }

    // Execute whois command with IP of server to get country
    command := "whois " + ip + " | grep -E '(C|c)ountry' | head -n 1 | sed -E 's/(C|c)ountry: *//'"
    country := utilities.RunShCommant(command)
    fmt.Printf(" - country: %s\n", country)

    // Execute whois command with IP of server to get owner
    command = "whois " + ip + " | grep -E 'OrgName|org-name|owner' | head -n 1 | sed -E 's/(OrgName|org-name|owner): *//'"
    owner := utilities.RunShCommant(command)
    fmt.Printf(" - owner: %s\n", owner)

    // Create and add server to servers
    server := &models.Server{0, ip, grade, country, owner, domain.ID}
    servers = append(servers, server)
  }

  return servers
}
