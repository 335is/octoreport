package octopus

import (
	"encoding/json"
)

// Feeds holds Octopus feeds
type Feeds struct {
	Items        []*Feed `json:"Items"`
	PagedResults PagedResults
}

// Feed describes a feed
type Feed struct {
	ID                                string   `json:"Id"`
	Name                              string   `json:"Name"`
	FeedType                          string   `json:"FeedType"`
	ApiVersion                        string   `json:"ApiVersion"`
	RegistryPath                      string   `json:"RegistryPath"`
	FeedUrl                           string   `json:"FeedUrl"`
	Username                          string   `json:"Username"`
	Password                          string   `json:"Password"`
	PackageAcquisitionLocationOptions []string `json:"PackageAcquisitionLocationOptions"`
	SpaceId                           string   `json:"SpaceId"`
	LastModifiedOn                    string   `json:"LastModifiedOn"`
	LastModifiedBy                    string   `json:"LastModifiedBy"`
}

// GetFeeds gets all Octopus feeds
func (c *Client) GetAllFeeds() (*[]Feed, error) {
	resp, err := c.DoGetRequest("feeds/all")
	if err != nil {
		return nil, err
	}

	feeds := []Feed{}
	err = json.NewDecoder(resp.Body).Decode(&feeds)
	if err != nil {
		return nil, err
	}

	return &feeds, nil
}
