package gogithub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// A Client represents a GitHub API client.
type Client struct {
	ID     string
	Secret string
}

// GetAccessToken calls access token API, gets an access token and returns it.
func (c *Client) GetAccessToken(code string) (string, error) {
	param := map[string]string{"client_id": c.ID, "client_secret": c.Secret, "code": code}
	b, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	res, err := http.Post(ACCESS_TOKEN_URL, "application/json", bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	m := map[string]string{}
	if err := parseResponse(res, &m); err != nil {
		return "", err
	}
	accessToken, prs := m["access_token"]
	if !prs {
		return "", fmt.Errorf("could not get an access token. [response: %+v]", m)
	}
	return accessToken, nil
}

// GetAuthenticatedUser gets the authenticated user and returns it.
func (c *Client) GetAuthenticatedUser(accessToken string) (*User, error) {
	res, err := http.Get(AUTHENTICATED_USER_URL + "?access_token=" + accessToken)
	if err != nil {
		return nil, err
	}
	u := &User{}
	if err := parseResponse(res, u); err != nil {
		return nil, err
	}
	return u, nil
}

// NewClient generates a client and returns it.
func NewClient(id string, secret string) *Client {
	return &Client{ID: id, Secret: secret}
}
