// Package handlers provides HTTP handlers for the API.
package handlers

import (
	"encoding/json"
	"net/http"
)

// Health responds with a basic JSON health check.
// It returns {"status":"ok"} with Content-Type application/json.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// Hello writes a plain text greeting.
// It returns 200 OK and "Hello, world!\n".
func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!\n"))
}

// Home writes a personal introduction message. Used on the root path.
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello sir, My name is Ian burns. I'll mostly choose the car rental system due to its complexity and usability. Also it would be useful for a family member who owns a car rental.\n"))
}

// About writes a short about paragraph.
func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("My name is Ian Burns, and I am an 18-year-old with a passion for cybersecurity. In my free time, I enjoy going on adventures and trying new things.\n"))
}

// Contact writes contact details.
func Contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contact me at:\nemail: ianburns344@gmail.com\nphone number: 6259908\n"))
}

// Qoute writes an inspirational quote. Note: kept the original spelling.
func Qoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("\u201CThe best way to get started is to quit talking and begin doing.\u201D – Walt Disney\n"))
}
