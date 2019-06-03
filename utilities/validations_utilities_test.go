package utilities

import (
	"testing"
)

func TestValidateDomainName(t *testing.T) {
	var domainName string

	domainName = "truora.com"
	if isValid, _ := ValidateDomainName(domainName); !isValid {
		t.Error("domain name is actually valid")
	}

	domainName = "SELECT * FROM domains;"
	if isValid, _ := ValidateDomainName(domainName); isValid {
		t.Error("domain name is actually invalid")
	}
}
