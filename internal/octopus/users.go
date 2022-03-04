package octopus

import (
	"encoding/json"
	"fmt"
)

// Users holds metrics on the deployment tasks
type Users struct {
	Items        []*User `json:"Items"`
	PagedResults PagedResults
}

// Task describes a deployment task, either running or not
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

// GetUsers gets list of Octopus users
func (c *Client) GetUsers() (*Users, error) {
	resp, err := c.DoGetRequest("users")
	if err != nil {
		return nil, err
	}

	users := Users{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (users *Users) Print() {
	for _, u := range users.Items {
		fmt.Println("Id             : ", u.ID)
		fmt.Println("Name           : ", u.Username)
	}
}
