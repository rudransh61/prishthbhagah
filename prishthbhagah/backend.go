// prishthbhagah.go
package prishthbhagah

import (
    "encoding/json"
    "net/http"
    "path/filepath"
    "strings"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, map[string]string)

type Router struct {
    routes map[string]map[string]HandlerFunc // Method -> Path -> HandlerFunc
}

func NewRouter() *Router {
    return &Router{routes: make(map[string]map[string]HandlerFunc)}
}

func (r *Router) Handle(method, path string, handler HandlerFunc) {
    if r.routes[method] == nil {
        r.routes[method] = make(map[string]HandlerFunc)
    }
    r.routes[method][path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    path := req.URL.Path
    method := req.Method

    for routePath, handler := range r.routes[method] {
        if routePath == path {
            handler(w, req, nil)
            return
        }

        if strings.Contains(routePath, ":") {
            parts := strings.Split(routePath, "/")
            reqParts := strings.Split(path, "/")

            if len(parts) != len(reqParts) {
                continue
            }

            params := make(map[string]string)
            match := true

            for i, part := range parts {
                if strings.HasPrefix(part, ":") {
                    paramName := strings.TrimPrefix(part, ":")
                    params[paramName] = reqParts[i]
                } else if part != reqParts[i] {
                    match = false
                    break
                }
            }

            if match {
                handler(w, req, params)
                return
            }
        }
    }

    http.NotFound(w, req)
}

func ServeFile(w http.ResponseWriter, req *http.Request, filePath string) {
    http.ServeFile(w, req, filepath.Clean(filePath))
}

func RespondJSON(w http.ResponseWriter, data interface{}, statusCode int) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}

func StartServer(router *Router, port string) error {
    return http.ListenAndServe(port, router)
}
