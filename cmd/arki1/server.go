package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	// Local port to listen to.
	Port = "8080"
	// Mocked user - should be detected using the app authentication database.
	User = "embeded-user@arki1.com"
)

func init() {
	flag.StringVar(&User, "user", "embeded-user@arki1.com",
		"The user `EMAIL` to be used as a authenticated mock")
	flag.StringVar(&Port, "port", "8080",
		"The `PORT` to listen to")
}

// main is the program entrypoint, and starts the HTTP server on Port.
func main() {
	flag.Parse()

	fmt.Println("Starting App Server for Looker Signed Embed Demo.")
	fmt.Println("Open http://localhost:" + Port)
	http.HandleFunc("/", index)
	http.HandleFunc("/dashboard/", dashboard)
	http.ListenAndServe(":"+Port, nil)
}

// index renders the index template showing the user a welcome message
// Here we should authenticate the user. For this demo we assume the user is already
// authorized.
func index(w http.ResponseWriter, r *http.Request) {
	if err := indexTpl.Execute(w, map[string]any{"user": User}); err != nil {
		log.Printf("Error rendering index: %v", err)
		return
	}
}

// dashboard renders the dashboard template and renders the embeded one from Looker.
// Before rendering the template, it will try to create an SSO embed URL using the Looker SDK.
func dashboard(w http.ResponseWriter, r *http.Request) {
	url, err := SignedEmbedURL(User)
	if err != nil {
		errorTpl.Execute(w, map[string]any{"error": err})
		return
	}
	if err := dashboardTpl.Execute(w, map[string]any{"url": url}); err != nil {
		log.Printf("Error rendering dashboard: %v", err)
	}
}
