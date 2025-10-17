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

type Checkpoint struct {
	Type      string `json:"type"`
	Data      string `json:"data"`
	PublicKey string `json:"public_key"`
}

type CheckpointResponse struct {
	AccountID  string     `json:"account_id"`
	Object     string     `json:"object"`
	Checkpoint Checkpoint `json:"checkpoint"`
}

type AccountStatus struct {
	Message     string `json:"message"`
	AccountID   string `json:"account_id"`
	AccountType string `json:"account_type"`
}
type WebhookResponse struct {
	AccountStatus AccountStatus
}

const (
	CheckpointStatusOTP             = "OTP"
	CheckpointStatusInAPPValidation = "IN_APP_VALIDATION"
	CheckpointStatusCAPTCHA         = "CAPTCHA"
	// More not implemented and handled currently
	WebhookMessageOK         = "OK"
	WebhookMessageConnecting = "CONNECTING"
	WebhookSyncSuccess       = "SYNC_SUCCESS"
)

type CheckpointOTPPayload struct {
	AccountID string `json:"account_id"`
	Code      string `json:"code"`
	Provider  string `json:"provider"`
}

type Account struct {
	Id        int64
	AccountId string
	Email     string
	Status    string
}
