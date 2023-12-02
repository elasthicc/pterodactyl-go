package pteroapp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Object     string `json:"object"`
	Attributes struct {
		ID         int    `json:"id"`
		ExternalID string `json:"external_id"`
		UUID       string `json:"uuid"`
		Username   string `json:"username"`
		Email      string `json:"email"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Language   string `json:"language"`
		RootAdmin  bool   `json:"root_admin"`
		TwoFactor  bool   `json:"2fa"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	} `json:"attributes"`
}

// UsersApplication is a client for the Users API
type UsersApplication struct {
	application *Application
}

// GetByID retrived the user by its ID.
func (a *UsersApplication) GetByID(ctx context.Context, id int64) (*User, *http.Response, error) {
	url := a.application.endpoint + fmt.Sprintf("users/%d", id)

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.application.token))

	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, res, fmt.Errorf("%s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}

	var userData User
	err = json.Unmarshal(body, &userData)
	if err != nil {
		return nil, nil, err
	}

	return &userData, req.Response, nil
}
