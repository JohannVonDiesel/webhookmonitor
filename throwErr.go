package webhookmonitor

import (
	"encoding/json"
	"net/http"
)

//-- any reponse ≠ 200 = err
func throwError(s string, w http.ResponseWriter) {
	http.Error(w, s, http.StatusNotFound)
}

func throwJSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
