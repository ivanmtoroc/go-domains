package models

import (
  "time"
  "testing"
)

func TestServerEqual(t *testing.T) {
  server_one := &Server{0, "34.23.15.1", "A+", "CO", "Truora", "truora.com", time.Now()}
  server_two := &Server{1, "34.23.15.1", "A+", "CO", "Truora", "truora.com", time.Now()}
  server_three := &Server{2, "34.22.15.1", "A+", "CO", "Truora", "truora.com", time.Now()}

  if equal := server_one.Equal(server_two); !equal {
    t.Error("Servers is actually equals")
  }

  if equal := server_one.Equal(server_three); equal {
    t.Error("Servers is actually differents")
  }
}
