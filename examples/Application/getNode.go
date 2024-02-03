package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/elasthicc/pterodactyl-go/pterogo/pteroapp"
)

func GetNodeByID() {

	url := "panel_url"
	token := "panel_token"
	myPteroApp := pteroapp.NewApplication(pteroapp.WithEndpoint(url), pteroapp.WithToken(token))

	// get node
	node, _, err := myPteroApp.Nodes.GetByID(context.Background(), 1)
	if err != nil {
		log.Fatalf("error retirving node: %s\n", err)
	}

	fmt.Printf("node name is: %s\n", node.Attributes.Name)

}
