package pteroapp

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestNodeAppCreate(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	env.Mux.HandleFunc("/api/application/nodes", func(w http.ResponseWriter, r *http.Request) {
		var reqBody NodeCreateOpts
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Fatal(err)
		}

		testNode := Node{}
		testNode.Attributes.Name = reqBody.Name
		json.NewEncoder(w).Encode(testNode)
	})

	nodeOpts := NodeCreateOpts{
		Name: "testNode",
	}
	node, _, err := env.Application.Nodes.Create(context.Background(), nodeOpts)
	if err != nil {
		t.Fatal(err)
	}

	if node.Attributes.Name != nodeOpts.Name {
		t.Errorf("unexpected node name: %v", node.Attributes.Name)
	}
}
