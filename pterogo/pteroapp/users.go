package pteroapp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// User represents a user in the Pterodactyl API.
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

type UsersList struct {
	Object string `json:"object"`
	Users  []User `json:"data"`
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

// UsersApplication is a client for the Users Pterodactyl API.
type UserApplication struct {
	application *Application
}

// GetByID retrives the user by its ID.
func (a *UserApplication) GetByID(ctx context.Context, id int64) (*User, *http.Response, error) {

	req, err := a.application.NewRequest(ctx, http.MethodGet, fmt.Sprintf("users/%d", id), nil)

	body, resp, err := a.application.Do(req)
	if err != nil {
		return nil, resp, err
	}

	var userData User
	err = json.Unmarshal(body, &userData)
	if err != nil {
		return nil, resp, err
	}

	return &userData, resp, nil
}

// GetUserList returns all of the users.
func (a *UserApplication) GetList(ctx context.Context) (*UsersList, *http.Response, error) {
	req, err := a.application.NewRequest(ctx, http.MethodGet, "users", nil)

	body, resp, err := a.application.Do(req)
	if err != nil {
		return nil, resp, err
	}

	var usersList UsersList
	err = json.Unmarshal(body, &usersList)
	if err != nil {
		return nil, resp, err
	}

	return &usersList, resp, nil
}

// UserCreateOpts specifies options for creating a new user.
type UserCreateOpts struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Create creates a new user.
func (a *UserApplication) Create(ctx context.Context, opts UserCreateOpts) (*User, *http.Response, error) {

	jsonReq, err := json.Marshal(opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := a.application.NewRequest(ctx, http.MethodPost, "users", bytes.NewBuffer(jsonReq))

	body, resp, err := a.application.Do(req)
	if err != nil {
		return nil, resp, err
	}

	var userResp User
	json.Unmarshal(body, &userResp)

	return &userResp, resp, nil
}
