package octopus

import (
	"encoding/json"
)

// Tenant describes a feed
type Tenant struct {
	ID                                string   `json:"Id"`
	Name                              string   `json:"Name"`
	TenantType                        string   `json:"TenantType"`
	ApiVersion                        string   `json:"ApiVersion"`
	RegistryPath                      string   `json:"RegistryPath"`
	TenantUrl                         string   `json:"TenantUrl"`
	Username                          string   `json:"Username"`
	Password                          string   `json:"Password"`
	PackageAcquisitionLocationOptions []string `json:"PackageAcquisitionLocationOptions"`
	SpaceId                           string   `json:"SpaceId"`
	LastModifiedOn                    string   `json:"LastModifiedOn"`
	LastModifiedBy                    string   `json:"LastModifiedBy"`
}

// GetTenants gets all Octopus tenants
func (c *Client) GetAllTenants() ([]Tenant, error) {
	tenants := []Tenant{}

	resp, err := c.DoGetRequest("tenants/all")
	if err != nil {
		return tenants, err
	}

	err = json.NewDecoder(resp.Body).Decode(&tenants)
	if err != nil {
		return tenants, err
	}

	return tenants, nil
}
