package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)
func TestHomeHandler(t *testing.T) {
    // set up request and response
    req := httptest.NewRequest("GET", "/", nil)
    rr := httptest.NewRecorder()
    // convert the home  function to a http.Handler type
    handler := http.HandlerFunc(home)
    // pass the fake HTTP request to the handler
    handler.ServeHTTP(rr, req)

// Cont'd on next slide
    //1. check the status code for 200 OK
    status := rr.Code
    if status != http.StatusOK {
      t.Errorf("got %v, expected %v", status, http.StatusOK)
    }
    //2. check the response body
    expected := "Hello from newsletter"
    got := rr.Body.String()
    if got != expected {
      t.Errorf("got %v, expected %v", got, expected)
    }
}

   
