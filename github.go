// Package gogithub implements a GitHub API client.
package gogithub

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// A Client represents a GitHub API client.
type Client struct {
	ID     string
	Secret string
}

// An AccessTokenResponse returns access token API's response data.
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// AccessToken calls access token API and returns the response data.
func (c *Client) AccessToken(code string) (int, *AccessTokenResponse, error) {
	param := map[string]string{"client_id": c.ID, "client_secret": c.Secret, "code": code}
	b, err := json.Marshal(param)
	if err != nil {
		return 0, nil, err
	}
	res, err := http.Post("https://github.com/login/oauth/access_token", "application/json", bytes.NewReader(b))
	if err != nil {
		return 0, nil, err
	}
	sCode := res.StatusCode
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return sCode, nil, err
	}
	var tokenRes AccessTokenResponse
	if err := json.Unmarshal(body, &tokenRes); err != nil {
		return sCode, nil, err
	}
	return sCode, &tokenRes, nil
}

// NewClient generates a client and returns it.
func NewClient(id string, secret string) *Client {
	return &Client{ID: id, Secret: secret}
}
