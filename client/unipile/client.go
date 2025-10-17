package unipile

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

var (
	Client IUnipileClient = UnipileClient{}
)

type IUnipileClient interface {
	ConnectUnipileAccount(jsonPayload []byte, unipileConfig UnipileConfig) *http.Response
}

type UnipileClient struct {
}

func (c UnipileClient) ConnectUnipileAccount(jsonPayload []byte, unipileConfig UnipileConfig) *http.Response {
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
	return resp
}
