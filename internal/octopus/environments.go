package octopus

import (
	"encoding/json"
	"fmt"
)

// Environments holds a list of Environments
type Environments struct {
	Items        []*Environment `json:"Items"`
	PagedResults PagedResults
}

// Environment is an octopus deployment environment, such as dev, test, prod
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

// GetEnvironments fetches the Octopus environments
func (c *Client) GetEnvironments() (*Environments, error) {
	resp, err := c.DoGetRequest("environments")
	if err != nil {
		return nil, err
	}

	envs := Environments{}
	err = json.NewDecoder(resp.Body).Decode(&envs)
	if err != nil {
		return nil, err
	}

	return &envs, nil
}

func (envs *Environments) Print() {
	for _, e := range envs.Items {
		fmt.Println("Id             : ", e.ID)
		fmt.Println("Space          : ", e.SpaceId)
		fmt.Println("Name           : ", e.Name)
	}
}
