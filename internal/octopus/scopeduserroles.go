package octopus

import (
	"encoding/json"
	"fmt"
)

// ScopedUserRoles is a list of ScopedUserRole
type ScopedUserRoles struct {
	Items []*ScopedUserRole `json:"Items"`
	PagedResults
}

// ScopedUserRole is the role that binds users and teams
type ScopedUserRole struct {
	ID              string   `json:"Id"`
	UserRoleID      string   `json:"UserRoleId"`
	TeamID          string   `json:"TeamId"`
	ProjectIDs      []string `json:"ProjectIds"`
	EnvironmentIDs  []string `json:"EnvironmentIds"`
	TenantIDs       []string `json:"TenantIds"`
	ProjectGroupIDs []string `json:"ProjectGroupIds"`
	SpaceID         string   `json:"SpaceId"`
}

// GetScopedUserRoles fetches the Octopus roles
func (c *Client) GetScopedUserRoles() (*ScopedUserRoles, error) {
	resp, err := c.DoGetRequest("scopeduserroles")
	if err != nil {
		return nil, err
	}

	roles := ScopedUserRoles{}
	err = json.NewDecoder(resp.Body).Decode(&roles)
	if err != nil {
		return nil, err
	}

	return &roles, nil
}

func (roles *ScopedUserRoles) Print() {
	for _, r := range roles.Items {
		fmt.Println("ID             : ", r.ID)
		fmt.Println("UserRoleID     : ", r.UserRoleID)
		fmt.Println("TeamID         : ", r.TeamID)
	}
}
