package pteroapp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Scheme represents the scheme for the node.
type Scheme string

// Constants for the supported schemes.
const (
	HTTP  Scheme = "http"
	HTTPS Scheme = "https"
)

// Node represents a node in the Pterodactyl API.
type Node struct {
	Object     string `json:"object"`
	Attributes struct {
		ID                 int    `json:"id"`
		UUID               string `json:"uuid"`
		Public             bool   `json:"public"`
		Name               string `json:"name"`
		Description        string `json:"description"`
		LocationID         int    `json:"location_id"`
		FQDN               string `json:"fqdn"`
		Scheme             string `json:"scheme"`
		BehindProxy        bool   `json:"behind_proxy"`
		MaintenanceMode    bool   `json:"maintenance_mode"`
		Memory             int    `json:"memory"`
		MemoryOverallocate int    `json:"memory_overallocate"`
		Disk               int    `json:"disk"`
		DiskOverallocate   int    `json:"disk_overallocate"`
		UploadSize         int    `json:"upload_size"`
		DaemonListen       int    `json:"daemon_listen"`
		DaemonSFTP         int    `json:"daemon_sftp"`
		DaemonBase         string `json:"daemon_base"`
		CreatedAt          string `json:"created_at"`
		UpdatedAt          string `json:"updated_at"`
	} `json:"attributes"`
}

// NodeApplication is a client for the Nodes API.
type NodeApplication struct {
	application *Application
}

// GetByID retrives the node by its ID.
func (a *NodeApplication) GetByID(ctx context.Context, id int64) (*Node, *http.Response, error) {
	url := a.application.endpoint + fmt.Sprintf("nodes/%d", id)

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

	var nodeData Node
	err = json.Unmarshal(body, &nodeData)
	if err != nil {
		return nil, nil, err
	}

	return &nodeData, resp, nil
}

// NodeCreateOpts represents the options for creating a new node.
type NodeCreateOpts struct {
	Name               string `json:"name"`
	LocationID         int    `json:"location_id"`
	FQDN               string `json:"fqdn"`
	Scheme             Scheme `json:"scheme"`
	Memory             int    `json:"memory"`
	MemoryOverallocate int    `json:"memory_overallocate"`
	Disk               int    `json:"disk"`
	DiskOverallocate   int    `json:"disk_overallocate"`
	UploadSize         int    `json:"upload_size"`
	DaemonSFTP         int    `json:"daemon_sftp"`
	DaemonListen       int    `json:"daemon_listen"`
}

// Create creates a new node.
func (a *NodeApplication) Create(ctx context.Context, opts NodeCreateOpts) (*Node, *http.Response, error) {

	jsonReq, err := json.Marshal(opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := a.application.NewRequest(ctx, http.MethodPost, "nodes", bytes.NewBuffer(jsonReq))

	body, resp, err := a.application.Do(req)
	if err != nil {
		return nil, resp, err
	}

	var nodeResp Node
	json.Unmarshal(body, &nodeResp)

	return &nodeResp, resp, nil

}
