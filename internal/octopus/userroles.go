package octopus

import (
	"encoding/json"
	"fmt"
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
func (c *Client) GetUserRoles() (*UserRoles, error) {
	resp, err := c.DoGetRequest("scopeduserroles")
	if err != nil {
		return nil, err
	}

	roles := UserRoles{}
	err = json.NewDecoder(resp.Body).Decode(&roles)
	if err != nil {
		return nil, err
	}

	return &roles, nil
}

func (roles *UserRoles) Print() {
	for _, r := range roles.Items {
		fmt.Println("Id             : ", r.ID)
		fmt.Println("Name           : ", r.Name)
		fmt.Println("Description    : ", r.Description)
	}
}
