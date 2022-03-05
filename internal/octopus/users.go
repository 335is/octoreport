package octopus

import (
	"encoding/json"
)

// Users holds metrics on the deployment tasks
type Users struct {
	Items        []*User `json:"Items"`
	PagedResults PagedResults
}

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
func (c *Client) GetAllUsers() (*[]User, error) {
	resp, err := c.DoGetRequest("users/all")
	if err != nil {
		return nil, err
	}

	users := []User{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}
