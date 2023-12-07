package pteroapp

import (
	"net/http"
	"net/http/httptest"
)

type testEnv struct {
	Server      *httptest.Server
	Mux         *http.ServeMux
	Application *Application
}

func (env *testEnv) Teardown() {
	env.Server.Close()
	env.Server = nil
	env.Mux = nil
	env.Application = nil
}

func newTestEnv() testEnv {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	application := NewApplication(
		WithEndpoint(server.URL),
		WithToken("token"),
	)
	return testEnv{
		Server:      server,
		Mux:         mux,
		Application: application,
	}
}
