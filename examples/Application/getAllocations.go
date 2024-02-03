package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/elasthicc/pterodactyl-go/pterogo/pteroapp"
)

func GetAllocationsByID() {

	url := "panel_url"
	token := "panel_token"
	myPteroApp := pteroapp.NewApplication(pteroapp.WithEndpoint(url), pteroapp.WithToken(token))

	// List allocations for a node
	allocations, _, err := myPteroApp.Nodes.ListAllocationsByID(context.Background(), 1)
	if err != nil {
		log.Fatalf("error retriving allocations: %s\n", err)
	}

	for _, allocation := range allocations.Data {
		fmt.Printf("ID: %d| Port:%d\n", allocation.Attributes.ID, allocation.Attributes.Port)
	}

}
