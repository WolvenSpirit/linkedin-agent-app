package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	unipile "github.com/wolvenspirit/linkedin-agent-app/client/unipile"
	models "github.com/wolvenspirit/linkedin-agent-app/models"
)

type UnipileMockClient struct {
	ConnectUnipileAccountResponse CheckpointResponse
}

var (
	ConnectUnipileAccountCalled = 0
)

func (c UnipileMockClient) ConnectUnipileAccount(jsonPayload []byte, unipileConfig unipile.UnipileConfig) *http.Response {
	rec := httptest.NewRecorder()
	b, _ := json.Marshal(c.ConnectUnipileAccountResponse)
	rec.Body = bytes.NewBuffer(b)
	ConnectUnipileAccountCalled++
	return rec.Result()
}

func prepare(resp CheckpointResponse) {
	clientMock := UnipileMockClient{}
	clientMock.ConnectUnipileAccountResponse = resp
	unipile.Client = clientMock
	models.Load()
}

func Test_loginLinkedin(t *testing.T) {
	w := httptest.NewRecorder()

	formData := url.Values{}
	formData.Set("username", "testuser")
	formData.Set("password", "testpassword")

	r := httptest.NewRequest(http.MethodPost, "/login/linkedin", bytes.NewBufferString(formData.Encode()))
	tests := []struct {
		name         string
		w            http.ResponseWriter
		r            *http.Request
		mockResponse CheckpointResponse
	}{
		{r: r, w: w, name: CheckpointStatusOTP, mockResponse: CheckpointResponse{AccountID: "abcde", Object: "checkpoint", Checkpoint: Checkpoint{Type: CheckpointStatusOTP}}},
		{r: r, w: w, name: CheckpointStatusInAPPValidation, mockResponse: CheckpointResponse{AccountID: "abcde", Object: "checkpoint", Checkpoint: Checkpoint{Type: CheckpointStatusInAPPValidation}}},
		{r: r, w: w, name: CheckpointStatusCAPTCHA, mockResponse: CheckpointResponse{AccountID: "abcde", Object: "checkpoint", Checkpoint: Checkpoint{Type: CheckpointStatusCAPTCHA}}},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare
			prepare(tt.mockResponse)
			mockdb, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			mock.ExpectExec(models.AccountModel.InsertAccount)
			db = mockdb

			// exec
			loginLinkedin(tt.w, tt.r)

			// assert
			if ConnectUnipileAccountCalled != 1+i {
				t.Error("Expected ConnectUnipileAccount to be called")
			}
		})

	}
}
