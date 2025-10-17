package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	models "github.com/wolvenspirit/linkedin-agent-app/models"
)

func main() {
	sigInt := make(chan os.Signal, 1)
	signal.Notify(sigInt, os.Interrupt)
	mux := http.NewServeMux()
	mux.HandleFunc("/login/linkedin", loginLinkedin)
	mux.HandleFunc("/webhook/accounts", webhookAccounts)
	mux.HandleFunc("/checkpoint/otp", checkpointOTPHandler)
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/status/linkedin", checkStatus)
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("port")),
		Handler: mux,
	}

	DBConnect()
	MigrateUp()
	models.Load()
	defer db.Close()
	// Start the server in a separate goroutine.
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-sigInt
	server.Close()
}
