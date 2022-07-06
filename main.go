package main

import (
	"log"
	"net/http"
	"os"
)

func getEnv(key string, defaultValue string) string {
	val, ok := os.LookupEnv(key)
	
	if !ok {
		return defaultValue
	} else {
		return val
	}
}

func main() {
	port := getEnv("PORT", "3000")
	redirectTo := getEnv("REDIRECT_To", "https://example.com")

	// handle any request
	http.HandleFunc("/", func (w http.ResponseWriter, r * http.Request) {
		// take the path and use the new domain instead
		redirectURI := redirectTo + r.RequestURI

		log.Printf("Redirected to %s", redirectURI)
		
		// redirect user
		http.Redirect(w, r, redirectURI, http.StatusPermanentRedirect)
	})

	log.Printf("Server started on port %v", port)
	// start server and log if error
	log.Fatal(http.ListenAndServe(":"+port, nil))
}