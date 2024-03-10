package pteroapp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Allocation represents an allocation in the Pterodactyl API.
type Allocation struct {
	Object     string `json:"object"`
	Attributes struct {
		ID       int     `json:"id"`
		IP       string  `json:"ip"`
		Alias    *string `json:"alias"`
		Port     int     `json:"port"`
		Notes    *string `json:"notes"`
		Assigned bool    `json:"assigned"`
	} `json:"attributes"`
}

// Allocations represents a list of allocations in the pterodactyl API.
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

// ListAllocationsByID lists the allocations on a node by ID.
func (a *NodeApplication) ListAllocationsByID(ctx context.Context, id int) (*Allocations, *http.Response, error) {
	req, err := a.application.NewRequest(ctx, http.MethodGet, fmt.Sprintf("nodes/%d/allocations", id), nil)

	body, resp, err := a.application.Do(req)
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

type AllocationCreateOpts struct {
	IP      string   `json:"ip"`
	Ports   []string `json:"ports"`
	IpAlias string   `json:"alias"`
}

func (a *NodeApplication) CreateAllocation(ctx context.Context, opts AllocationCreateOpts, nodeID int) (*http.Response, error) {

	jsonReq, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	req, err := a.application.NewRequest(ctx, http.MethodPost, fmt.Sprintf("nodes/%d/allocations", nodeID), bytes.NewBuffer(jsonReq))

	if err != nil {
		return nil, err
	}

	_, resp, err := a.application.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
