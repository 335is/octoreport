package octopus

import (
	"encoding/json"
)

// Channels holds octopus channels
type Channels struct {
	Items []Channel `json:"Items"`
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
//	note: channels does not support the /all route (it does in newer Octopus versions)
func (c *Client) GetAllChannels() ([]Channel, error) {
	channels := []Channel{}

	resp, err := c.DoGetRequest("channels")
	if err != nil {
		return channels, err
	}

	chInfo := Channels{}
	err = json.NewDecoder(resp.Body).Decode(&chInfo)
	if err != nil {
		return channels, err
	}

	ch := make([]Channel, len(chInfo.Items))
	copy(ch, chInfo.Items)

	return ch, nil
}
