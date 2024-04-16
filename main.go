// main.go
package main

import (
	"fmt"
	"prishthbhagah/prishthbhagah"
	"net/http"
)

func main() {
	app := prishthbhagah.NewApp()

	// GET request to render an HTML file
	app.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		prishthbhagah.RenderHTML(w, "hello.html", nil)
	})

	// GET request to respond with JSON
	app.Get("/api/data", func(w http.ResponseWriter, r *http.Request) {
		prishthbhagah.JSON(w, map[string]string{"message": "Data received successfully"})
	})

	// GET request with parameters
	app.Get("/user/:id/:name", func(w http.ResponseWriter, r *http.Request) {
        // Print URL path for debugging
        fmt.Println("URL Path:", r.URL.Path)
    
        // Print parameters for debugging
        params := prishthbhagah.Params(r)
        fmt.Println("Parameters:", params)
    
        // Access parameters without colon prefix
        userID := params["id"]
        userName := params["name"]
        prishthbhagah.JSON(w, map[string]string{"userID": userID, "userName": userName})
    })
    

	// Start the server
	port := "8080"
	if err := app.Start(port); err != nil {
		panic(err)
	}

	// Print server link
	fmt.Printf("Server is listening on http://localhost:%s\n", port)
}
