package scraper

import (
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

  log.Println("Consulting API | Host name: " + domain.Name)
  for {
    // Make http request (Method Get) to API
    response, err := http.Get(API_URL + domain.Name)
    if err != nil {
      log.Println("API request error")
      log.Println(err)
      return nil, err
    }

    // Decode response
    err = json.NewDecoder(response.Body).Decode(&result)
    if err == nil {
      // Validate status of domain
      status := result["status"].(string)
      log.Println("  - Status: " + status)
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

//
func getServerFromResponse(domain *models.Domain, result map[string]interface{}) []*models.Server {
  var servers []*models.Server

  // By default domain have the best SSL grade and is down
  domain.SslGrade = "A+"
  domain.IsDown = true

  log.Println("  - Get endpoints")

  // Iterate endpoints of JSON response
  for i, element := range result["endpoints"].([]interface{}) {
    // Type assertion to map
    endpoint := element.(map[string]interface{})
    log.Printf("     - Endpoint %d:\n", i)

    // Get IP of server
    ip := endpoint["ipAddress"].(string)
    log.Printf("        - IP: %s\n", ip)

    // By default all servers have invalid SSL
    grade := "Unknown"
    status := endpoint["statusMessage"].(string)
    // If status is Ready get SSL grade
    if status == "Ready" {
      grade = endpoint["grade"].(string)
      // If at least one server is up then domain is not down
      domain.IsDown = false
    }
    log.Printf("        - Grade: %s\n", grade)
    // Update min SSL Grade of domain
    if GRADES[grade] > GRADES[domain.SslGrade] {
      domain.SslGrade = grade
    }

    // Execute whois command with IP of server to get country
    command := "whois " + ip + " | grep -E '(C|c)ountry' | head -n 1 | sed -E 's/(C|c)ountry: *//'"
    country := utilities.RunShCommant(command)
    log.Printf("        - Country: %s\n", country)

    // Execute whois command with IP of server to get owner
    command = "whois " + ip + " | grep -E 'OrgName|org-name|owner' | head -n 1 | sed -E 's/(OrgName|org-name|owner): *//'"
    owner := utilities.RunShCommant(command)
    log.Printf("        - Owner: %s\n", owner)

    // Create and add server to servers
    server := &models.Server{0, ip, grade, country, owner, domain.ID}
    servers = append(servers, server)
  }

  if domain.SslGrade == "Unknown" {
    domain.SslGrade = ""
  }

  return servers
}
