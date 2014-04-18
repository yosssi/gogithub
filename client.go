package gogithub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// A Client represents a GitHub API client.
type Client struct {
	ID          string
	Secret      string
	AccessToken string
}

// SetAccessToken calls access token API, gets an access token and sets it to the client.
func (c *Client) SetAccessToken(code string) error {
	param := map[string]string{"client_id": c.ID, "client_secret": c.Secret, "code": code}
	b, err := json.Marshal(param)
	if err != nil {
		return err
	}
	res, err := http.Post(AccessTokenURL, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	m := map[string]string{}
	if err := parseResponse(res, &m); err != nil {
		return err
	}
	accessToken, prs := m[ParamKeyAccessToken]
	if !prs {
		return fmt.Errorf("could not get an access token. [response: %+v]", m)
	}
	c.AccessToken = accessToken
	return nil
}

// GetAuthenticatedUser gets the authenticated user and returns it.
func (c *Client) GetAuthenticatedUser() (*User, error) {
	if c.AccessToken == "" {
		return nil, fmt.Errorf("access token is not set to the client.")
	}
	res, err := http.Get(AuthenticatedUserURL + c.AccessTokenURLParam())
	if err != nil {
		return nil, err
	}
	u := &User{}
	if err := parseResponse(res, u); err != nil {
		return nil, err
	}
	return u, nil
}

// GetContents gets the contents of the specified path.
func (c *Client) GetContents(owner string, repo string, path string) (*Contents, error) {
	res, err := http.Get(fmt.Sprintf(GetContentsPath, owner, repo, path) + c.AccessTokenURLParam())
	if err != nil {
		return nil, err
	}
	contents := &Contents{}
	if err := parseResponse(res, contents); err != nil {
		return nil, err
	}
	return contents, nil
}

// AccessTokenURLParam returns the access token url parameter.
func (c *Client) AccessTokenURLParam() string {
	if c.AccessToken == "" {
		return ""
	}
	return URLParamPrefix + ParamKeyAccessToken + "=" + c.AccessToken
}

// NewClient generates a client and returns it.
func NewClient(id string, secret string) *Client {
	return &Client{ID: id, Secret: secret}
}
