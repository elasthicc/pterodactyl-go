package pteroapp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Servers represents the servers in the Pterodactyl panel
type Server struct {
	Object     string `json:"object"`
	Attributes struct {
		ID            int                 `json:"id"`
		ExternalID    string              `json:"external_id"`
		UUID          string              `json:"uuid"`
		Identifier    string              `json:"identifier"`
		Name          string              `json:"name"`
		Description   string              `json:"description"`
		Suspended     bool                `json:"suspended"`
		Limits        ServerLimits        `json:"limits"`
		FeatureLimits ServerFeatureLimits `json:"feature_limits"`
		User          int                 `json:"user"`
		Node          int                 `json:"node"`
		Allocation    int                 `json:"allocation"`
		Nest          int                 `json:"nest"`
		Egg           int                 `json:"egg"`
		Pack          interface{}         `json:"pack"`
		Container     ServerContainer     `json:"container"`
		UpdatedAt     string              `json:"updated_at"`
		CreatedAt     string              `json:"created_at"`
	} `json:"attributes"`
}

// ServerContainer represents the container information for a server.
type ServerContainer struct {
	StartupCommand string            `json:"startup_command"`
	Image          string            `json:"image"`
	Installed      int               `json:"installed"`
	Environment    ServerEnvironment `json:"environment"`
}

// ServerLimits represents the resource limits for a server.
type ServerLimits struct {
	Memory int `json:"memory"`
	Swap   int `json:"swap"`
	Disk   int `json:"disk"`
	IO     int `json:"io"`
	CPU    int `json:"cpu"`
}

// ServerFeatureLimits represents the feature limits for a server.
type ServerFeatureLimits struct {
	Databases int `json:"databases"`
	Backups   int `json:"backups"`
}

// ServerAllocation represents the allocation for a server.
type ServerAllocation struct {
	Default int `json:"default"`
}

// ServerEnvironment represents the environment variables for a server.
type ServerEnvironment struct {
	MinecraftVersion       string `json:"MINECRAFT_VERSION"`
	ServerJarFile          string `json:"SERVER_JARFILE"`
	DLPath                 string `json:"DL_PATH"`
	BuildNumber            string `json:"BUILD_NUMBER"`
	Startup                string `json:"STARTUP"`
	PServerLocation        string `json:"P_SERVER_LOCATION"`
	PServerUUID            string `json:"P_SERVER_UUID"`
	PServerAllocationLimit int    `json:"P_SERVER_ALLOCATION_LIMIT"`
}

// ServersApplication is a client for the Servers API
type ServerApplication struct {
	application *Application
}

// GetByID retrives the server by its ID.
func (a *ServerApplication) GetByID(ctx context.Context, id int64) (*Server, *http.Response, error) {
	url := a.application.endpoint + fmt.Sprintf("servers/%d", id)

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

	var serverData Server

	err = json.Unmarshal(body, &serverData)
	if err != nil {
		return nil, nil, err
	}

	return &serverData, resp, nil
}

// ServerCreateOpts represents the options for creating a new server.
type ServerCreateOpts struct {
	Name          string              `json:"name"`
	User          int                 `json:"user"`
	LocationID    int                 `json:"location_id"`
	Node          int                 `json:"node"` //
	Nest          int                 `json:"nest"` //
	Egg           int                 `json:"egg"`
	DockerImage   string              `json:"docker_image"`
	Startup       string              `json:"startup"`
	Environment   ServerEnvironment   `json:"environment"`
	Limits        ServerLimits        `json:"limits"`
	FeatureLimits ServerFeatureLimits `json:"feature_limits"`
	Allocation    ServerAllocation    `json:"allocation"`
}

// Create creates a server
func (a *ServerApplication) Create(opts ServerCreateOpts) (Server, *http.Response, error) {
	url := a.application.endpoint + "servers"
	method := "POST"

	jsonReq, err := json.Marshal(opts)
	if err != nil {
		return Server{}, nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))

	if err != nil {
		return Server{}, nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.application.token))

	resp, err := client.Do(req)
	if err != nil {
		return Server{}, nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Server{}, nil, err
	}

	if resp.StatusCode != 201 {
		return Server{}, resp, fmt.Errorf("%s", resp.Status)
	}

	var serverResp Server
	json.Unmarshal(bodyBytes, &serverResp)

	return serverResp, resp, nil
}
