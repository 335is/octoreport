package octopus

import (
	"encoding/json"
	"time"
)

// Tasks holds metrics on the deployment tasks
type Tasks struct {
	ItemType       string
	TotalResults   int
	ItemsPerPage   int
	NumberOfPages  int
	LastPageNumber int
	TotalCounts    *TotalCounts
	Items          []Task
	Links          Links
}

// TotalCounts holds metrics on currently running deployment tasks
type TotalCounts struct {
	Canceled    int
	Cancelling  int
	Executing   int
	Failed      int
	Queued      int
	Success     int
	TimedOut    int
	Interrupted int
}

// Task describes a deployment task, either running or not
type Task struct {
	ID                         string `json:"Id"`
	SpaceID                    string `json:"SpaceId"`
	Name                       string
	Description                string
	Arguments                  map[string]string
	State                      string
	Completed                  time.Time
	QueueTime                  time.Time
	QueueTimeExpiry            time.Time
	StartTime                  time.Time
	LastUpdatedTime            time.Time
	CompletedTime              time.Time
	ServerNode                 string
	Duration                   string
	ErrorMessage               string
	HasBeenPickedUpByProcessor bool
	IsCompleted                bool
	FinishedSuccessfully       bool
	HasPendingInterruptions    bool
	CanRerun                   bool
	HasWarningsOrErrors        bool
	Links                      Links
}

// GetStats gets metrics on all tasks
func (c *Client) GetStats() (*Tasks, error) {
	resp, err := c.DoGetRequest("tasks")
	if err != nil {
		return nil, err
	}

	tasks := Tasks{}
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return &tasks, nil
}

// GetAll gets all tasks
func (c *Client) GetAll() (*[]Task, error) {
	resp, err := c.DoGetRequest("tasks")
	if err != nil {
		return nil, err
	}

	tasks := Tasks{}
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return &tasks.Items, nil
}

// GetFiltered gets all tasks matched by the specified query parameters
func (c *Client) GetFiltered(qp map[string]string) ([]Task, error) {
	resp, err := c.DoGetRequestParms("tasks", qp)
	if err != nil {
		return nil, err
	}

	tasks := Tasks{}
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks.Items, nil
}

// GetByID gets a task by its ID
func (c *Client) GetByID(ID string) (*Task, error) {
	resp, err := c.DoGetRequest("tasks" + ID)
	if err != nil {
		return nil, err
	}

	task := Task{}
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
