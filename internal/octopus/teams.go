package octopus

import (
	"encoding/json"
)

// Teams holds octopus teams
type Teams struct {
	Items []*Team `json:"Items"`
	PagedResults
}

// Team describes a team of users
type Team struct {
	ID            string   `json:"Id"`
	Name          string   `json:"Name"`
	MemberUserIds []string `json:"MemberUserIds"`
	// ExternalSecurityGroups []string `json:"ExternalSecurityGroups,omitempty"`
	UserRoleIds      []string `json:"UserRoleIds"`
	ProjectIds       []string `json:"ProjectIds"`
	EnvironmentIds   []string `json:"EnvironmentIds"`
	TenantIds        []string `json:"TenantIds"`
	TenantTags       []string `json:"TenantTags"`
	ProjectGroupIds  []string `json:"ProjectGroupIds"`
	CanBeDeleted     bool     `json:"CanBeDeleted"`
	CanBeRenamed     bool     `json:"CanBeRenamed"`
	CanChangeRoles   bool     `json:"CanChangeRoles"`
	CanChangeMembers bool     `json:"CanChangeMembers"`
}

// GetAllTeams fetches all Octopus teams and related data
func (c *Client) GetAllTeams() (*[]Team, error) {
	resp, err := c.DoGetRequest("teams/all")
	if err != nil {
		return nil, err
	}

	teams := []Team{}
	err = json.NewDecoder(resp.Body).Decode(&teams)
	if err != nil {
		return nil, err
	}

	return &teams, nil
}
