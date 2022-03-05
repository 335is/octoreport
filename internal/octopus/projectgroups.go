package octopus

import (
	"encoding/json"
)

// ProjectGroups holds the project groups
type ProjectGroups struct {
	Items        []*ProjectGroup `json:"Items"`
	PagedResults PagedResults
}

// ProjectGroup describes a project group
type ProjectGroup struct {
	ID                string   `json:"Id"`
	Name              string   `json:"Name"`
	Description       string   `json:"Description"`
	EnvironmentIds    []string `json:"EnvironmentIds"`
	RetentionPolicyId string   `json:"RetentionPolicyId"`
}

// GetProjectGroups gets all Octopus project groups
func (c *Client) GetAllProjectGroups() (*[]ProjectGroup, error) {
	resp, err := c.DoGetRequest("projectgroups/all")
	if err != nil {
		return nil, err
	}

	groups := []ProjectGroup{}
	err = json.NewDecoder(resp.Body).Decode(&groups)
	if err != nil {
		return nil, err
	}

	return &groups, nil
}
