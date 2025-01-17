package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	// Local port to listen to.
	Port = flag.String("port",
		"8080",
		"The `PORT` to listen to")

	// Mocked user - should be detected using the app authentication database.
	User = flag.String("user",
		env("DEMO_USER", "external-user@example.com"),
		"The user `EMAIL` to be used as a authenticated mocked user")

	// DashboardURL is the URL to the item that we want to embed in our app.
	// Could be a Look or any other embeddable item URL.
	DashboardURL = flag.String("dashboard-url",
		env("DEMO_DASHBOARD", "https://arki1.cloud.looker.com/embed/dashboards/17"),
		"The dashboard `URL` to be embedded")
)

// main is the program entrypoint, and starts the HTTP server on Port.
func main() {
	flag.Parse()

	fmt.Println("Starting App Server for Looker Signed Embed Demo.")
	fmt.Println("Open http://localhost:" + *Port)

	http.HandleFunc("/", index)
	http.HandleFunc("/dashboard/", dashboard)
	http.ListenAndServe(":"+*Port, nil)
}

// index renders the index template showing the user a welcome message
//
// *Note*: Here we should authenticate the user. For this demo we assume the user is already
// authorized.
func index(w http.ResponseWriter, r *http.Request) {
	if err := indexTpl.Execute(w, map[string]any{"user": User}); err != nil {
		log.Printf("Error rendering index: %v", err)
		return
	}
}

// dashboard renders the dashboard template and renders the embeded one from Looker.
// Before rendering the template, it will try to create an SSO embed URL using the Looker SDK.
//
// *Note*: Here we should authenticate the user. For this demo we assume the user is already
// authorized.
func dashboard(w http.ResponseWriter, r *http.Request) {
	url, err := SignedEmbedURL(*User)
	if err != nil {
		errorTpl.Execute(w, map[string]any{"error": err})
		return
	}
	if err := dashboardTpl.Execute(w, map[string]any{"url": url}); err != nil {
		log.Printf("Error rendering dashboard: %v", err)
	}
}

// env is a helper function to get an environment variable with a fallback
// value if it is not defined.
func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
