package scraper

import (
	"fmt"
	"log"
	"time"
	"bytes"
	"os/exec"
	"strings"
	"net/http"
	"encoding/json"
	"golang.org/x/net/html"
	"go-domains/models"
)

var GRADES = map[string]int {
	"A+":	1,
	"A":	2,
	"B":	3,
	"C":	4,
	"D":	5,
	"E":	6,
	"F":	7,
}

func GetDomainAPI(hostName string) (*models.Domain, []*models.Server) {
	domain := &models.Domain{hostName, false, "", "", getIcon(hostName), false, time.Now()}
	servers := getServers(domain)

	if servers == nil {
		domain.IsDown = true
	}

	return domain, servers
}

func getServers(domain *models.Domain) []*models.Server {
	const API_URL = "https://api.ssllabs.com/api/v3/analyze?host="
	var result map[string]interface{}
	var servers []*models.Server
	num_min_grade := 1

	for {
		response, err := http.Get(API_URL + domain.Name)
		if err != nil {
			return nil
		}

		json.NewDecoder(response.Body).Decode(&result)

		fmt.Println(result["status"])
		if status := result["status"]; status == "READY" {
			break
		} else if status == "ERROR" {
			return nil
		}
		time.Sleep(1 * time.Second)
	}

	for _, endpoint := range result["endpoints"].([]interface{}) {
		ip := endpoint.(map[string]interface{})["ipAddress"].(string)
		mensaje := endpoint.(map[string]interface{})["statusMessage"].(string)
		grade := "Certificate not valid."
		if mensaje == "Ready" {
			grade = endpoint.(map[string]interface{})["grade"].(string)
		}
		country := runShCommant("whois " + ip + " | grep -E '[Cc]ountry' | head -n 1 | sed -E 's/[Cc]ountry: *//'")
		owner := runShCommant("whois " + ip + " | grep -E 'OrgName|org-name|owner' | head -n 1 | sed -E 's/(OrgName|org-name|owner): *//'")
		server := &models.Server{ip, grade, country, owner, domain.Name, time.Now()}
		servers = append(servers, server)
		if num_min_grade < GRADES[grade] {
			num_min_grade = GRADES[grade]
		}
	}

	for k, v := range GRADES {
		if v == num_min_grade {
			domain.SslGrade = k
		}
	}

	return servers
}

func runShCommant(command string) string {
	cmd := exec.Command("/bin/sh", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(out.String(), "\n")
}

func getIcon(hostName string) string {
	url := "http://" + hostName
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body := response.Body
	defer body.Close()

	doc := html.NewTokenizer(body)

	for {
		tt := doc.Next()
		switch tt {
			case html.SelfClosingTagToken:
				token := doc.Token()
				if token.Data == "link" {
					for _, link := range token.Attr {
						if link.Key == "rel" && strings.Contains(link.Val, "icon") {
							for _, link := range token.Attr {
								if link.Key == "href" {
									if string(link.Val[0]) == "/" {
										return url + link.Val
									}
									return link.Val
								}
							}
							break
						}
					}
				}
			case html.ErrorToken:
				return "Icon not found."
		}
	}
}
