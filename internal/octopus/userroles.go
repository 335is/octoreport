package octopus

import (
	"encoding/json"
)

// UserRoles is a list of UserRoles
type UserRoles struct {
	Items []*UserRole `json:"Items"`
	PagedResults
}

// UserRole is the role that binds users and teams
type UserRole struct {
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// GetUserRoles fetches the Octopus roles
func (c *Client) GetAllUserRoles() (*[]UserRole, error) {
	resp, err := c.DoGetRequest("userroles/all")
	if err != nil {
		return nil, err
	}

	roles := []UserRole{}
	err = json.NewDecoder(resp.Body).Decode(&roles)
	if err != nil {
		return nil, err
	}

	return &roles, nil
}
