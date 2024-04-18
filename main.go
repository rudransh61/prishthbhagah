package main

import (
    "fmt"
    "prishthbhagah/prishthbhagah"
    "net/http"
)

func main() {
    router := prishthbhagah.NewRouter()

    // Serve the index.html file
    router.Handle("GET", "/", func(w http.ResponseWriter, req *http.Request, _ map[string]string) {
        prishthbhagah.ServeFile(w, req, "./static/index.html")
    })

    // Serve the form.html file
    router.Handle("GET", "/form", func(w http.ResponseWriter, req *http.Request, _ map[string]string) {
        prishthbhagah.ServeFile(w, req, "./static/form.html")
    })

    // Handle POST requests to /submit
    router.Handle("POST", "/submit", func(w http.ResponseWriter, req *http.Request, _ map[string]string) {
        err := req.ParseForm()
        if err != nil {
            http.Error(w, "Failed to parse form", http.StatusBadRequest)
            return
        }
        name := req.Form.Get("name")
        fmt.Fprintf(w, "Hello, %s! Your form was submitted successfully.", name)
    })

    // Handle GET requests with parameters
    router.Handle("GET", "/hello/:name", func(w http.ResponseWriter, req *http.Request, params map[string]string) {
        name := params["name"]
        fmt.Fprintf(w, "Hello, %s!", name)
    })

    // Start server
    if err := prishthbhagah.StartServer(router, ":8080"); err != nil {
        panic(err)
    }
}
