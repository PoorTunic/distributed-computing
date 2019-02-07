package web

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// ResponseHeaderContentTypeKey set the content type
	ResponseHeaderContentTypeKey = "Content-Type"
	// ResponseHeaderContentTypeJSONUTF8 set the application and charset
	ResponseHeaderContentTypeJSONUTF8 = "application/json; charset=UTF-8"
)

// ParamAsString returns an URL parameter /{name} as a string
func ParamAsString(name string, r *http.Request) string {
	vars := mux.Vars(r)
	value := vars[name]
	return value
}

// GetJSONContent returns the JSON content of a request
func GetJSONContent(v interface{}, r *http.Request) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

// SendJSONWithHTTPCode outputs JSON with an HTTP code
func SendJSONWithHTTPCode(w http.ResponseWriter, d interface{}, code int) {
	w.Header().Set(ResponseHeaderContentTypeKey, ResponseHeaderContentTypeJSONUTF8)
	w.WriteHeader(code)
	if d != nil {
		err := json.NewEncoder(w).Encode(d)
		if err != nil {
			panic(err)
		}
	}
}

// SendJSONOk outputs a JSON with http.StatusOK code
func SendJSONOk(w http.ResponseWriter, d interface{}) {
	SendJSONWithHTTPCode(w, d, http.StatusOK)
}

// SendJSONError sends error with a custom message and error code
func SendJSONError(w http.ResponseWriter, error string, code int) {
	SendJSONWithHTTPCode(w, map[string]string{"errorMsg": error}, code)
}

// SendJSONNotFound produces a http.StatusNotFound response
func SendJSONNotFound(w http.ResponseWriter) {
	SendJSONError(w, "Resource not found", http.StatusNotFound)
}

// SendJSONOkNotAction produces a http.StatusNotFound response
func SendJSONOkNotAction(w http.ResponseWriter) {
	SendJSONWithHTTPCode(w, "Everything was deleted", http.StatusOK)
}
