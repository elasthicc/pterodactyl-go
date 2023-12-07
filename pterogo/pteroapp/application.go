package pteroapp

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Application struct {
	endpoint   string
	token      string
	httpClient *http.Client

	Users UserApplication
	Nodes NodeApplication
	//Locations Locations
	Servers ServerApplication
	//Nests     Nests
}

type ApplicationOption func(*Application)

func WithEndpoint(endpoint string) ApplicationOption {
	return func(app *Application) {
		app.endpoint = strings.TrimRight(endpoint, "/") + "/api/application/"
	}
}

func WithToken(token string) ApplicationOption {
	return func(app *Application) {
		app.token = token
	}
}

func NewApplication(options ...ApplicationOption) *Application {
	app := &Application{}

	for _, option := range options {
		option(app)
	}

	app.httpClient = &http.Client{}

	app.Users = UserApplication{
		application: app,
	}
	app.Nodes = NodeApplication{
		application: app,
	}
	app.Servers = ServerApplication{
		application: app,
	}

	return app
}

// NewRequest creates an HTTP request against the API, the returned request
// is assigned with ctx and has all necessary headers (auth, content type, authorization)
func (a *Application) NewRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	url := a.endpoint + path

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	if a.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.token))
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req = req.WithContext(ctx)

	return req, nil
}

// Do performs an HTTP request against the API.
func (a *Application) Do(r *http.Request) ([]byte, *http.Response, error) {

	resp, err := a.httpClient.Do(r)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		return nil, resp, err
	}

	if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
		err = fmt.Errorf("pterogo: server responded with status code %d", resp.StatusCode)
		return body, resp, err
	}

	return body, resp, nil
}
