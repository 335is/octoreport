package octopus

import (
	"encoding/json"
)

// Lifecycles holds lifecycles
type Lifecycles struct {
	Items        []*Lifecycle `json:"Items"`
	PagedResults PagedResults
}

// Lifecycle describes a lifecycle
type Lifecycle struct {
	ID                      string          `json:"Id"`
	Name                    string          `json:"Name"`
	Description             string          `json:"Description"`
	SpaceId                 string          `json:"SpaceId"`
	Phases                  []Phase         `json:"Phases"`
	ReleaseRetentionPolicy  RetentionPolicy `json:"ReleaseRetentionPolicy"`
	TentacleRetentionPolicy RetentionPolicy `json:"TentacleRetentionPolicy"`
}

// Phase describes a lifecycle phase
type Phase struct {
	ID                                 string   `json:"Id"`
	Name                               string   `json:"Name"`
	AutomaticDeploymentTargets         []string `json:"AutomaticDeploymentTargets"`
	OptionalDeploymentTargets          []string `json:"OptionalDeploymentTargets"`
	MinimumEnvironmentsBeforePromotion int      `json:"MinimumEnvironmentsBeforePromotion"`
	IsOptionalPhase                    bool     `json:"IsOptionalPhase"`
}

type RetentionPolicy struct {
	Unit              string `json:"Unit"`
	QuantityToKeep    int    `json:"QuantityToKeep"`
	ShouldKeepForever bool   `json:"ShouldKeepForever"`
}

// GetLifecycles gets all Octopus lifecycles
func (c *Client) GetAllLifecycles() (*[]Lifecycle, error) {
	resp, err := c.DoGetRequest("lifecycles/all")
	if err != nil {
		return nil, err
	}

	lifecycles := []Lifecycle{}
	err = json.NewDecoder(resp.Body).Decode(&lifecycles)
	if err != nil {
		return nil, err
	}

	return &lifecycles, nil
}
