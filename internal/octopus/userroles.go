package octopus

import (
	"encoding/json"
)

// UserRole is the role that binds users and teams
type UserRole struct {
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

// GetUserRoles fetches the Octopus roles
func (c *Client) GetAllUserRoles() ([]UserRole, error) {
	userRoles := []UserRole{}

	resp, err := c.DoGetRequest("userroles/all")
	if err != nil {
		return userRoles, err
	}

	err = json.NewDecoder(resp.Body).Decode(&userRoles)
	if err != nil {
		return userRoles, err
	}

	return userRoles, nil
}
