package octopus

import (
	"encoding/json"
)

// Project describes a project
type Project struct {
	ID                  string `json:"Id"`
	Name                string `json:"Name"`
	Slug                string `json:"Slug"`
	Description         string `json:"Description"`
	VariableSetID       string `json:"VariableSetId"`
	DeploymentProcessID string `json:"DeploymentProcessId"`
	ProjectGroupId      string `json:"ProjectGroupId"`
	LifecycleId         string `json:"LifecycleId"`
}

// GetAllProjects gets all Octopus projects
func (c *Client) GetAllProjects() ([]Project, error) {
	projects := []Project{}

	resp, err := c.DoGetRequest("projects/all")
	if err != nil {
		return projects, err
	}

	err = json.NewDecoder(resp.Body).Decode(&projects)
	if err != nil {
		return projects, err
	}

	return projects, nil
}
