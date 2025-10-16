package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func loginLinkedin(w http.ResponseWriter, r *http.Request) {
	unipileConfig := GetUnipileConfig()

	// Create the request payload.
	payload := ConnectAccountPayload{
		Provider: unipileConfig.Provider,
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	// Marshal the payload into JSON.
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
	}

	// Create a new HTTP POST request.
	url := fmt.Sprintf("https://%s/api/v1/accounts", unipileConfig.UnipileDsn)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Error creating request: %v", err)
	}

	// Set the required headers.
	req.Header.Set("X-API-KEY", unipileConfig.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// Send the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
	}

	// Check for a successful response.
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Printf("API request failed with status code %d and response: %s", resp.StatusCode, string(body))
	}

	// Unmarshal the JSON response into our struct.
	//var accountResponse AccountResponse
	checkpoint := CheckpointResponse{}
	if err := json.Unmarshal(body, &checkpoint); err != nil {
		log.Printf("Error unmarshaling response JSON: %v", err)
	}
	log.Printf("Response %s", checkpoint.Checkpoint.Type)
	if checkpoint.Checkpoint.Type == CheckpointStatusOTP {
		log.Printf("Checkpoint OTP required for account ID %s", checkpoint.AccountID)
		templates := template.Must(template.ParseFiles("templates/otp_checkpoint.html"))
		data := make(map[string]interface{}, 1)
		data["AccountID"] = checkpoint.AccountID
		if err := templates.Execute(w, data); err != nil {
			log.Print(err.Error())
		}
	}
	if checkpoint.Checkpoint.Type == CheckpointStatusInAPPValidation {
		log.Printf("Checkpoint In App Validation required for account ID %s", checkpoint.AccountID)
		templates := template.Must(template.ParseFiles("templates/in_app_validation_screen.html"))
		data := make(map[string]interface{}, 1)
		if err := templates.Execute(w, data); err != nil {
			log.Print(err.Error())
		}
	}
	//w.WriteHeader(http.StatusOK)
}

func webhookAccounts(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
	}
	result := WebhookResponse{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Error unmarshaling response JSON: %v", err)
	}
	log.Printf("Received webhook: %+v", result)
	if result.Message == WebhookMessageOK || result.Message == WebhookSyncSuccess {
		// save to db as below
		// query row by account_id
		// update row with connection status as CONNECTED
	}
	if result.Message == WebhookMessageConnecting {
		// update row with connection status as CONNECTING
	}
	w.WriteHeader(http.StatusOK)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		templates := template.Must(template.ParseFiles("templates/login_form.html"))
		data := make(map[string]interface{}, 1)
		if err := templates.Execute(w, data); err != nil {
			log.Print(err.Error())
		}
	}
}

func checkpointOTPHandler(w http.ResponseWriter, r *http.Request) {

	unipileConfig := GetUnipileConfig()

	// Create the request payload.
	payload := CheckpointOTPPayload{
		AccountID: r.FormValue("account_id"),
		Code:      r.FormValue("otp"),
		Provider:  unipileConfig.Provider,
	}

	// Marshal the payload into JSON.
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
	}

	// Create a new HTTP POST request.
	url := fmt.Sprintf("https://%s/api/v1/accounts/checkpoint", unipileConfig.UnipileDsn)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("Error creating request: %v", err)
	}

	// Set the required headers.
	req.Header.Set("X-API-KEY", unipileConfig.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	// Send the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read and print the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
	}

	// Check for a successful response.
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		log.Printf("API request failed with status code %d and response: %s", resp.StatusCode, string(body))
	}

	log.Printf("Response %s", string( // Create a new HTTP POST request.
		body))
	w.WriteHeader(http.StatusOK)

}
