package pteroapp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestServerAppGetByID(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	fmt.Println(env.Application.endpoint)

	env.Mux.HandleFunc("/api/application/servers/1", func(w http.ResponseWriter, r *http.Request) {
		testServer := Server{}
		testServer.Attributes.ID = 1
		json.NewEncoder(w).Encode(testServer)
	})

	ctx := context.Background()

	user, _, err := env.Application.Servers.GetByID(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}

	if user.Attributes.ID != 1 {
		t.Errorf("unexpected user ID: %v", user.Attributes.ID)
	}
}

func TestServerAppCreate(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	env.Mux.HandleFunc("/api/application/servers", func(w http.ResponseWriter, r *http.Request) {
		var reqBody ServerCreateOpts
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Fatal(err)
		}

		testServer := Server{}
		testServer.Attributes.Name = reqBody.Name
		json.NewEncoder(w).Encode(testServer)
	})

	serverOpts := ServerCreateOpts{
		Name: "testServer",
	}
	server, _, err := env.Application.Servers.Create(context.Background(), serverOpts)
	if err != nil {
		t.Fatal(err)
	}

	if server.Attributes.Name != serverOpts.Name {
		t.Errorf("unexpected user ID: %v", server.Attributes.Name)
	}
}
