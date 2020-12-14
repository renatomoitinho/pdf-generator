package converts

import (
    "encoding/base64"
    "fmt"
    "net/http"
    "pdf-generation/internal/pkg/constants"
)

func ConverterHandler(w http.ResponseWriter, r *http.Request) {
    sources := r.URL.Query().Get("src")

    if len(sources) == 0 {
        unprocessableEntity(&w, r, "param 'src' is required.")
        return
    }

    url, err := decodeBase64(sources)
    if err != nil {
        unprocessableEntity(&w, r, "parameter 'src' does not contain the base64 value of the content.")
        return
    }

    cr := &ConvertRequest{Url: url}

    buffer, err := cr.ToPdf()

    if err != nil {
        unprocessableEntity(&w, r, err.Error())
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set(constants.HeadersContentType, constants.ContentTypeJson)
    _, _ = w.Write(buffer)

}

func decodeBase64(sources string) (string, error) {
    v, err := base64.StdEncoding.DecodeString(sources)
    if err != nil {
        return "", err
    }
    return string(v), nil
}

func unprocessableEntity(w *http.ResponseWriter, r *http.Request, message string) {
    str := fmt.Sprintf(`{"status": %d,"message": "%s", "path": "%s" }`, http.StatusUnprocessableEntity, message, r.RequestURI)
    writer := *w
    writer.Header().Set(constants.HeadersContentType, constants.ContentTypeJson)
    writer.WriteHeader(http.StatusUnprocessableEntity)
    _, _ = writer.Write([]byte(str))
}
