package main

import (
	"net/http"
	"testing"
)

func TestNoSurg(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:	
		// do nothing
		default:
			t.Errorf("type is not Http.Handler, but is %t", v)
	}

}

func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:	
		// do nothing
		default:
			t.Errorf("type is not Http.Handler, but is %t", v)
	}

}