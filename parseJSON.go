package webhookmonitor

import (
	"encoding/json"
	"log"
	"net/http"
)

//---- Struct ---- //
type webhookJSONRespStruct struct {
	Alert struct {
		ClearValue  int    `json:"clear_value"`
		AlertValue  string `json:"alert_value"`
		CirconusRef string `json:"circonus_ref"`
		Agent       string `json:"agent"`
		ClearTime   string `json:"clear_time"`
		CheckName   string `json:"check_name"`
		AlertID     string `json:"alert_id"`
		Severity    int    `json:"severity"`
		AlertTime   string `json:"alert_time"`
		AlertURL    string `json:"alert_url"`
		MetricName  string `json:"metric_name"`
	} `json:"alertBody"`
	AccountName string `json:"account_name"`
}

//-- function decodes JSON response from webhook
func parseJSON(r *http.Request, w http.ResponseWriter) bool {
	log.Println("Parsed JSON")
	//-- decode JSON from request body
	decoder := json.NewDecoder(r.Body)
	//-- t instantiates as webhookJSONRespStruct struct
	var t webhookJSONRespStruct
	//-- catch err
	err := decoder.Decode(&t)
	if err != nil {
		log.Println("Error: ", err)
		return false
	}
	//-- log incoming webhook
	log.Println("Incoming webhook from: ", t.AccountName)
	return true
}
