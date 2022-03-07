package octopus

import (
	"encoding/json"
)

// ProjectGroup describes a project group
type ProjectGroup struct {
	ID                string   `json:"Id"`
	Name              string   `json:"Name"`
	Description       string   `json:"Description"`
	EnvironmentIds    []string `json:"EnvironmentIds"`
	RetentionPolicyId string   `json:"RetentionPolicyId"`
}

// GetProjectGroups gets all Octopus project groups
func (c *Client) GetAllProjectGroups() ([]ProjectGroup, error) {
	projectGroups := []ProjectGroup{}

	resp, err := c.DoGetRequest("projectgroups/all")
	if err != nil {
		return projectGroups, err
	}

	err = json.NewDecoder(resp.Body).Decode(&projectGroups)
	if err != nil {
		return projectGroups, err
	}

	return projectGroups, nil
}
