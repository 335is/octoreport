package octopus

import (
	"encoding/json"
	"fmt"
)

// root describes the JSON response from the root API call
type root struct {
	Application    string `json:"Application"`
	Version        string `json:"Version"`
	APIVersion     string `json:"ApiVersion"`
	InstallationID string `json:"InstallationId"`
	Links          map[string]string
}

// Version health checks the root API and gets the Octopus Deploy version
func (c *Client) Version() (string, error) {
	response, err := c.DoGetRequest("")
	if err != nil {
		return "Failed GET request", err
	}

	root := root{}
	err = json.NewDecoder(response.Body).Decode(&root)
	if err != nil {
		return fmt.Sprintf("Failed to decode JSON response from %s: %v", c.Address, err), err
	}

	return fmt.Sprintf("%s %s %s %s", root.Application, root.Version, root.APIVersion, root.InstallationID), nil
}
