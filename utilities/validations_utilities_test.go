package utilities

import (
  "testing"
)

func TestValidateDomainName(t *testing.T) {
  var domain_name string

  domain_name = "truora.com"
  if is_valid := ValidateDomainName(domain_name); !is_valid {
    t.Error("Domain name is actually valid")
  }

  domain_name = "SELECT * FROM domains;"
  if is_valid := ValidateDomainName(domain_name); is_valid {
    t.Error("Domain name is actually invalid")
  }
}
