package octopus

import (
	"encoding/json"
)

// User describes a deployment task, either running or not
type User struct {
	ID                  string `json:"Id"`
	Username            string `json:"Username"`
	DisplayName         string `json:"DisplayName"`
	IsActive            bool   `json:"IsActive"`
	IsService           bool   `json:"IsService"`
	EmailAddress        string `json:"EmailAddress"`
	CanPasswordBeEdited bool   `json:"CanPasswordBeEdited"`
	IsRequestor         bool   `json:"IsRequestor"`
}

// GetUsers gets all Octopus users
func (c *Client) GetAllUsers() ([]User, error) {
	users := []User{}

	resp, err := c.DoGetRequest("users/all")
	if err != nil {
		return users, err
	}

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return users, err
	}

	return users, nil
}
