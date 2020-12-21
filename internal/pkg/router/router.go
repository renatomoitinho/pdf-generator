package router

import (
    "github.com/go-chi/chi"
    "github.com/go-chi/cors"
    "net/http"
    "pdf-generation/internal/pkg/constants"

    "pdf-generation/internal/app/converts"
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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set(constants.HeadersContentType, constants.ContentTypeJson)
    _, _ = w.Write([]byte(`{ "is_alive": true }`))
}

func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
    w.WriteHeader(http.StatusNotFound)
}

func MethodNotAllowedHandler(w http.ResponseWriter, _ *http.Request) {
    w.WriteHeader(http.StatusMethodNotAllowed)
}
