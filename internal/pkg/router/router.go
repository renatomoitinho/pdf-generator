package router

import (
    "encoding/json"
    "fmt"
    "github.com/go-chi/chi"
    "github.com/go-chi/cors"
    "net/http"
    "pdf-generation/internal/app/converts"
    "time"
)

func Router() (*chi.Mux, error) {
    router := chi.NewRouter()
    router.Use(cors.Handler(cors.Options{
       AllowedOrigins: []string{"*"},
       AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
       AllowedHeaders: []string{"*"},
       AllowCredentials: true,
    }))
    router.NotFound(NotFoundHandler)
    router.MethodNotAllowed(MethodNotAllowedHandler)
    router.Get("/", IndexHandler)
    router.Get("/convert", converts.ConverterHandler)

    return router, nil
}

func IndexHandler(w http.ResponseWriter, _ *http.Request) {

    response, err := json.Marshal(map[string] interface{} {
        "alive" : true,
        "now": time.Now().Format(time.RFC3339),
    })
    if err != nil {
        _ = fmt.Errorf(err.Error())
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    _, err = w.Write(response)
    if err != nil {
        _ = fmt.Errorf(err.Error())
    }
}

func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
    w.WriteHeader(http.StatusNotFound)
}

func MethodNotAllowedHandler(w http.ResponseWriter, _ *http.Request) {
    w.WriteHeader(http.StatusMethodNotAllowed)
}
