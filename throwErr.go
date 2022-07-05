package webhookmonitor

import (
	"net/http"
)

//-- any reponse ≠ 200 = err
func throwError(s string, w http.ResponseWriter) {
	http.Error(w, s, http.StatusNotFound)
	//--- TODO: output JSON err message
}
