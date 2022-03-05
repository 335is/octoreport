package octopus

import (
	"encoding/json"
)

// Projects holds the projects
type Projects struct {
	Items        []*Project `json:"Items"`
	PagedResults PagedResults
}

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
func (c *Client) GetAllProjects() (*[]Project, error) {
	resp, err := c.DoGetRequest("projects/all")
	if err != nil {
		return nil, err
	}

	projects := []Project{}
	err = json.NewDecoder(resp.Body).Decode(&projects)
	if err != nil {
		return nil, err
	}

	return &projects, nil
}
