package octopus

import (
	"encoding/json"
)

// Channels holds octopus channels
type Channels struct {
	Items []*Channel `json:"Items"`
	PagedResults
}

// Channel describes a channel
type Channel struct {
	ID          string   `json:"Id"`
	Name        string   `json:"Name"`
	Description string   `json:"Description"`
	ProjectId   string   `json:"ProjectId"`
	LifecycleId string   `json:"LifecycleId"`
	IsDefault   bool     `json:"IsDefault"`
	Rules       []string `json:"Rules"`
	TenantTags  []string `json:"TenantTags"`
	SpaceId     string   `json:"SpaceId"`
}

// GetAllChannels fetches all Octopus channels and related data
func (c *Client) GetAllChannels() (*[]Channel, error) {
	resp, err := c.DoGetRequest("channels/all")
	if err != nil {
		return nil, err
	}

	channels := []Channel{}
	err = json.NewDecoder(resp.Body).Decode(&channels)
	if err != nil {
		return nil, err
	}

	return &channels, nil
}
