package pteroapp

import (
	"strings"
)

type Application struct {
	endpoint string
	token    string

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
