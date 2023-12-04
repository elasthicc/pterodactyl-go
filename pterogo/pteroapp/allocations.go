package pteroapp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Allocation represents an allocation in the API response.
type Allocation struct {
	Object     string `json:"object"`
	Attributes struct {
		ID       int         `json:"id"`
		IP       string      `json:"ip"`
		Alias    interface{} `json:"alias"`
		Port     int         `json:"port"`
		Notes    interface{} `json:"notes"`
		Assigned bool        `json:"assigned"`
	} `json:"attributes"`
}

// Allocations represents a list of allocations in the API response.
type Allocations struct {
	Object string       `json:"object"`
	Data   []Allocation `json:"data"`
	Meta   struct {
		Pagination struct {
			Total       int      `json:"total"`
			Count       int      `json:"count"`
			PerPage     int      `json:"per_page"`
			CurrentPage int      `json:"current_page"`
			TotalPages  int      `json:"total_pages"`
			Links       struct{} `json:"links"`
		} `json:"pagination"`
	} `json:"meta"`
}

// ListAllocationsByID lists the allocations on a node by ID
func (a *NodeApplication) ListAllocationsByID(id int) (*Allocations, *http.Response, error) {
	url := a.application.endpoint + fmt.Sprintf("nodes/%d/allocations", id)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.application.token))

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, resp, fmt.Errorf("%s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	var allocations Allocations
	err = json.Unmarshal(body, &allocations)
	if err != nil {
		return nil, nil, err
	}

	return &allocations, resp, nil
}
