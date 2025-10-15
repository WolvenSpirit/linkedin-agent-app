package main

// Define the struct for the request body.
type ConnectAccountPayload struct {
	Provider string `json:"provider"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Define the struct for the response body.
type AccountResponse struct {
	AccountID string `json:"account_id"`
	Status    string `json:"status"`
	// Add other fields from the API response as needed.
}

type CheckpointResponse struct {
	AccountID  string `json:"account_id"`
	Object     string `json:"object"`
	Checkpoint struct {
		Type      string `json:"type"`
		Data      string `json:"data"`
		PublicKey string `json:"public_key"`
	} `json:"checkpoint"`
}

const (
	CheckpointStatusOTP             = "OTP"
	CheckpointStatusInAPPValidation = "IN_APP_VALIDATION"
	// More not implemented and handled currently
)

type CheckpointOTPPayload struct {
	AccountID string `json:"account_id"`
	Code      string `json:"code"`
	Provider  string `json:"provider"`
}
