package main_test

import (
	"net/http"
	"testing"
)

func TestGetOriginalURL(t *testing.T) {
	// make a dummy request
	resp, err := http.Get("http://localhost:1205/v1/short/1")

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, resp.StatusCode)
	}

	if err != nil {
		t.Errorf("Encountered an error:%s\n", err)
	}
}
