package pteroapp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestNodeAppListAllocationsByID(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	fmt.Println(env.Application.endpoint)

	env.Mux.HandleFunc("/api/application/nodes/1/allocations", func(w http.ResponseWriter, r *http.Request) {
		testAllocations := Allocations{}
		testAllocationOne := Allocation{}
		testAllocationTwo := Allocation{}

		testAllocationOne.Attributes.ID = 1
		testAllocationTwo.Attributes.ID = 2

		allocationArr := []Allocation{testAllocationOne, testAllocationTwo}

		testAllocations.Data = allocationArr

		json.NewEncoder(w).Encode(testAllocations)
	})

	ctx := context.Background()

	allocations, _, err := env.Application.Nodes.ListAllocationsByID(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}

	for i, allocation := range allocations.Data {
		if allocation.Attributes.ID != i+1 {
			t.Errorf("unexpected user ID: %v", allocation.Attributes.ID)
		}
	}
}
