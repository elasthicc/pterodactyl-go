package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/elasthicc/pterodactyl-go/pterogo/pteroapp"
)

func CreateAllocation() {

	url := "panel_url"
	token := "panel_token"
	myPteroApp := pteroapp.NewApplication(pteroapp.WithEndpoint(url), pteroapp.WithToken(token))

	// Create an allocation for a node

	myAllocations := pteroapp.AllocationCreateOpts{
		IP:    "127.0.0.0", // This must reflect your node's IP
		Ports: []string{"25565", "25571", "25572"},
	}

	resp, err := myPteroApp.Nodes.CreateAllocation(context.Background(), myAllocations, 1)
	if err != nil {
		log.Fatalf("error retriving allocations: %s\n", err)
	}

	fmt.Printf("request status: %s\n", resp.Status)
}
