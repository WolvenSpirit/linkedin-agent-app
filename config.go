package main

import "os"

type UnipileConfig struct {
	UnipileDsn  string
	AccessToken string
	Provider    string
}

func GetUnipileConfig() UnipileConfig {
	UnipileDsn := os.Getenv("unipile_dsn")
	AccessToken := os.Getenv("unipile_access_token")
	Provider := "LINKEDIN"
	return UnipileConfig{UnipileDsn, AccessToken, Provider}
}
