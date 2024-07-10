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
	Port = "8080"

	// Mocked user - should be detected using the app authentication database.
	User = "embeded-user@arki1.com"

	// DashboardURL is the URL to the item that we want to embed in our app.
	// Could be a Look or any other embeddable item URL.
	DashboardURL = "https://arki1.cloud.looker.com/embed/dashboards/17"
)

func init() {
	// Setup some command line options and flags
	flag.StringVar(&Port, "port",
		env("PORT", "8080"),
		"The `PORT` to listen to")
	flag.StringVar(&User, "user",
		env("DEMO_USER", "embeded-user@arki1.com"),
		"The user `EMAIL` to be used as a authenticated mocked user")
	flag.StringVar(&DashboardURL, "dashboard-url",
		env("DEMO_DASHBOARD", "https://arki1.cloud.looker.com/embed/dashboards/17"),
		"The dashboard `URL` to be embedded")
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

func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
