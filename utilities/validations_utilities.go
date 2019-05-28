package utilities

import (
  "log"
  "regexp"
)

func ValidateDomainName(name string) bool {
  re := `^[a-zA-Z0-9-.]*$`
  matched, err := regexp.MatchString(re, name)
  if err != nil {
    log.Println("regular expression error")
    log.Fatalln("- error: ", err)
  }
  return matched
}
