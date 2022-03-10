package octopus

import (
	"encoding/json"
)

// Machine describes a machine
type Machine struct {
	ID                string   `json:"Id"`
	Name              string   `json:"Name"`
	Thumbprint        string   `json:"Thumbprint"`
	URI               string   `json:"Uri"`
	IsDisabled        bool     `json:"IsDisabled"`
	EnvironmentIDs    []string `json:"EnvironmentIds"`
	Roles             []string `json:"Roles"`
	MachinePolicyId   string   `json:"MachinePolicyId"`
	TenantIDs         []string `json:"TenantIds"`
	TenantTags        []string `json:"TenantTags"`
	Status            string   `json:"Status"`
	HealthStatus      string   `json:"HealthStatus"`
	HasLatestCalamari bool     `json:"HasLatestCalamari"`
	StatusSummary     string   `json:"StatusSummary"`
	IsInProcess       bool     `json:"IsInProcess"`
	Endpoint          Endpoint `json:"Endpoint"`
}

type Endpoint struct {
	CommunicationStyle            string                 `json:"CommunicationStyle"`
	URI                           string                 `json:"Uri"`
	ProxyID                       string                 `json:"ProxyID"`
	Thumbprint                    string                 `json:"Thumbprint"`
	TentacleVersionDetails        TentacleVersionDetails `json:"TentacleVersionDetails"`
	CertificateSignatureAlgorithm string                 `json:"CertificateSignatureAlgorithm"`
	ID                            string                 `json:"Id"`
	LastModifiedOn                string                 `json:"LastModifiedOn"`
	LastModifiedBy                string                 `json:"LastModifiedBy"`
}

type TentacleVersionDetails struct {
	UpgradeLocked    bool   `json:"UpgradeLocked"`
	Version          string `json:"Version"`
	UpgradeSuggested bool   `json:"UpgradeSuggested"`
	UpgradeRequired  bool   `json:"UpgradeRequired"`
}

// GetMachines gets all Octopus machines
func (c *Client) GetAllMachines() ([]Machine, error) {
	machines := []Machine{}

	resp, err := c.DoGetRequest("machines/all")
	if err != nil {
		return machines, err
	}

	err = json.NewDecoder(resp.Body).Decode(&machines)
	if err != nil {
		return machines, err
	}

	return machines, nil
}
