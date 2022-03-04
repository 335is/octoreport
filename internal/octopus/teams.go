package octopus

import (
	"encoding/json"
	"fmt"
)

// Teams holds data on the octopus teams
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
	CanBeDeleted     bool   `json:"CanBeDeleted"`
	CanBeRenamed     bool   `json:"CanBeRenamed"`
	CanChangeRoles   bool   `json:"CanChangeRoles"`
	CanChangeMembers bool   `json:"CanChangeMembers"`
	SpaceId          string `json:"SpaceId"`
	Description      string `json:"Description"`
}

// GetTeams fetches the Octopus teams and related data
func (c *Client) GetTeams() (*Teams, error) {
	resp, err := c.DoGetRequest("teams")
	if err != nil {
		return nil, err
	}

	teams := Teams{}
	err = json.NewDecoder(resp.Body).Decode(&teams)
	if err != nil {
		return nil, err
	}

	return &teams, nil
}

func (teams *Teams) Print() {
	for _, t := range teams.Items {
		fmt.Println("Name           : ", t.Name)
		fmt.Println("Id             : ", t.ID)
		fmt.Println("Members        : ", FormatStringList(t.MemberUserIds))
	}
}
