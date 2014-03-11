// Package gogithub implements a GitHub API client.
package gogithub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// A Client represents a GitHub API client.
type Client struct {
	ID          string
	Secret      string
	AccessToken string
}

// An AccessTokenResponse returns access token API's response data.
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// SetAccessToken calls access token API, gets an access token and set it to the client.
func (c *Client) SetAccessToken(code string) error {
	param := map[string]string{"client_id": c.ID, "client_secret": c.Secret, "code": code}
	b, err := json.Marshal(param)
	if err != nil {
		return err
	}
	res, err := http.Post("https://github.com/login/oauth/access_token", "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	m := make(map[string]string)
	if err := json.Unmarshal(body, &m); err != nil {
		return err
	}
	accessToken, prs := m["access_token"]
	if !prs {
		return fmt.Errorf("could not get an access token. [response: %s]", string(body))
	}
	c.AccessToken = accessToken
	return nil
}

// NewClient generates a client and returns it.
func NewClient(id string, secret string) *Client {
	return &Client{ID: id, Secret: secret}
}
