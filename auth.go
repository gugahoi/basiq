package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

const ServerURL = "https://au-api.basiq.io/"

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func getAuthToken(apikey string) string {
	payload := strings.NewReader("scope=SERVER_ACCESS")
	req, _ := http.NewRequest("POST", ServerURL+"token", payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Basic "+apikey)
	req.Header.Add("basiq-version", "3.0")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("unable to authenticate", err)
	}

	defer res.Body.Close()
	raw, _ := io.ReadAll(res.Body)
	body := string(raw)

	if res.StatusCode != 200 {
		log.Fatalln("unable to authenticate", body)
	}

	var authResponse AuthResponse
	err = json.Unmarshal(raw, &authResponse)
	if err != nil {
		log.Fatalln("unable to parse access_token", body)
	}

	return authResponse.AccessToken
}
