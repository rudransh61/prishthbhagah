// myframework.go
package prishthbhagah

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	// "path/filepath"
	"strings"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type App struct {
	routes map[string]map[string]HandlerFunc
}

func NewApp() *App {
	return &App{
		routes: make(map[string]map[string]HandlerFunc),
	}
}

func (app *App) Get(path string, handler HandlerFunc) {
	if app.routes[path] == nil {
		app.routes[path] = make(map[string]HandlerFunc)
	}
	app.routes[path]["GET"] = handler
}

func (app *App) Post(path string, handler HandlerFunc) {
	if app.routes[path] == nil {
		app.routes[path] = make(map[string]HandlerFunc)
	}
	app.routes[path]["POST"] = handler
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path
	if strings.HasSuffix(urlPath, "/") {
		urlPath = urlPath[:len(urlPath)-1]
	}

	for route, handlers := range app.routes {
		if strings.HasPrefix(urlPath, route) {
			handler := handlers[r.Method]
			if handler != nil {
				handler(w, r)
				return
			}
		}
	}
	http.NotFound(w, r)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func RenderHTML(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func JSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *App) Start(port string) error {
	http.Handle("/", Logger(app))
	return http.ListenAndServe(":"+port, nil)
}

func Params(r *http.Request) map[string]string {
	params := make(map[string]string)
	parts := strings.Split(r.URL.Path, "/")

	for i := 0; i < len(parts); i++ {
		if strings.HasPrefix(parts[i], ":") && i+1 < len(parts) {
			paramName := strings.TrimPrefix(parts[i], ":")
			params[paramName] = parts[i+1]
		}
	}

	return params
}
