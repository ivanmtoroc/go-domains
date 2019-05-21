package models

import (
  "time"
  "testing"
)

func TestServerEqual(test *testing.T) {
  server_one := &Server{"34.23.15.1", "A+", "CO", "Truora", "truora.com", time.Now()}
  server_two := &Server{"34.23.15.1", "A+", "CO", "Truora", "truora.com", time.Now()}
  server_three := &Server{"34.22.15.1", "A+", "CO", "Truora", "truora.com", time.Now()}

  if equal := server_one.Equal(server_two); !equal {
    test.Error("Servers is actually equals")
  }

  if equal := server_one.Equal(server_three); equal {
    test.Error("Servers is actually differents")
  }
}
