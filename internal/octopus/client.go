package octopus

import (
	"net/http"
	"strings"
)

// Client is an Octopus Deploy API client
type Client struct {
	Address    string
	APIKey     string
	httpClient *http.Client
}

// New creates a new Octopus Deploy API client
func New(address, apiKey string, httpClient *http.Client) *Client {
	client := Client{
		Address:    strings.TrimRight(address, "/"),
		APIKey:     apiKey,
		httpClient: httpClient,
	}

	return &client
}

// PagedResults is used in most API calls
type PagedResults struct {
	ItemType       string `json:"ItemType"`
	TotalResults   int    `json:"TotalResults"`
	ItemsPerPage   int    `json:"ItemsPerPage"`
	NumberOfPages  int    `json:"NumberOfPages"`
	LastPageNumber int    `json:"LastPageNumber"`
	IsStale        bool   `json:"IsStale"`
	Links          Links  `json:"Links"`
}

// Links are part of each JSON response, used for pagination
type Links map[string]string

// FormatStringList prints a slice of strings
func FormatStringList(list []string) string {
	return "{" + strings.Join(list[:], ",") + "}"
}
