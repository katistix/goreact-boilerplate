package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
)

//go:embed frontend/dist/*
var frontend embed.FS

// Serve React static files from the embedded filesystem
func serveFrontend() http.Handler {
	// Open the embedded frontend directory
	subFS, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		fmt.Printf("Error accessing embedded files: %v\n", err)
		os.Exit(1)
	}
	fs := http.FileServer(http.FS(subFS))
	return http.StripPrefix("/", fs)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the health check logic here
	fmt.Fprintln(w, "Server is healthy")
}

func main() {
	// Handle API routes
	http.HandleFunc("/api/health", healthCheckHandler)

	// Serve the React frontend for all other routes
	http.Handle("/", serveFrontend())

	// Start the server on port 8080
	fmt.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
