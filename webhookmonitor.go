package webhookmonitor

import (
	"log"
	"net/http"
)

// InboundWebhook is an HTTP Cloud Function with a request parameter.
func InboundWebhook(w http.ResponseWriter, r *http.Request) {
	//-- log request
	log.Println("Incomming Request")

	//-- attempt to decode JSON from webbhook
	boolJSONParse := parseJSON(r, w)

	//-- throw err if JSON is shitty
	if !boolJSONParse {
		throwError("Unable to parse JSON response", w)
		return
	}
	//-- feel the joy
	w.Write([]byte("Success parsing incoming webhook JSON"))
}
