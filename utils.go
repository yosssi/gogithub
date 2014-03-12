package gogithub

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// parseResponse parses the HTTP response.
func parseResponse(res *http.Response, v interface{}) error {
	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}
	return nil
}
