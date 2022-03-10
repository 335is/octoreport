package octopus

import (
	"encoding/json"
)

// MachineRole holds a machine role
type MachineRole string

// GetMachines gets all Octopus machines
func (c *Client) GetAllMachineRoles() ([]MachineRole, error) {
	machineroles := []MachineRole{}

	resp, err := c.DoGetRequest("machineroles/all")
	if err != nil {
		return machineroles, err
	}

	err = json.NewDecoder(resp.Body).Decode(&machineroles)
	if err != nil {
		return machineroles, err
	}

	return machineroles, nil
}
