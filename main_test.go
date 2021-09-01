package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_addCounter(t *testing.T) {

	tests := []struct {
		name string
		req  *http.Request
	}{
		{
			name: "increment",
			req:  httptest.NewRequest("PUT", "/api/counter", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getCounter)
			handler.ServeHTTP(rr, tt.req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

		})
	}
}

func Test_getCounter(t *testing.T) {

	tests := []struct {
		name string
		req  *http.Request
	}{
		{
			name: "getCounter",
			req:  httptest.NewRequest("GET", "/api/counter", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getCounter)
			handler.ServeHTTP(rr, tt.req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})
	}
}

func Test_subCounter(t *testing.T) {

	tests := []struct {
		name string
		req  *http.Request
	}{
		{
			name: "decrement",
			req:  httptest.NewRequest("DELETE", "/api/counter", nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(getCounter)
			handler.ServeHTTP(rr, tt.req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
		})
	}
}
