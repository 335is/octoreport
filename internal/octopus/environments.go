package octopus

import (
	"encoding/json"
)

// Environment describes an octopus deployment environment, such as dev, test, prod
type Environment struct {
	ID                         string `json:"Id"`
	SpaceId                    string
	Name                       string
	Description                string
	SortOrder                  int
	UseGuidedFailure           bool
	AllowDynamicInfrastructure bool
	ExtensionSettings          []Extension
	Links                      Links
}

// Extension defines a particular Octopus extension, such as Jira
type Extension struct {
	ExtensionId string
	Values      map[string]string
}

// GetAllEnvironments fetches the Octopus environments
func (c *Client) GetAllEnvironments() ([]Environment, error) {
	environments := []Environment{}
	resp, err := c.DoGetRequest("environments/all")
	if err != nil {
		return environments, err
	}

	err = json.NewDecoder(resp.Body).Decode(&environments)
	if err != nil {
		return environments, err
	}

	return environments, nil
}
