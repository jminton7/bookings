package main

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/jminton7/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Errorf("type is not *chi.Mux, but is %t", v)
	}
}