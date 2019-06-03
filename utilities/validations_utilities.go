package utilities

import (
	"log"
	"regexp"
)

// ValidateDomainName test the domain name with a regex
func ValidateDomainName(name string) (bool, error) {
	re := "^[a-zA-Z0-9-.]*$"
	matched, err := regexp.MatchString(re, name)
	if err != nil {
		log.Println("regular expression error")
		log.Println("- error: ", err)
		return false, err
	}
	return matched, nil
}
