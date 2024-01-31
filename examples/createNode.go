package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/elasthicc/pterodactyl-go/pterogo/pteroapp"
)

func CreateNode() {

	url := "panel_url"
	token := "panel_token"
	myPteroApp := pteroapp.NewApplication(pteroapp.WithEndpoint(url), pteroapp.WithToken(token))

	// Create a node
	nodeOpts := pteroapp.NodeCreateOpts{
		Name:               "pterogo",
		LocationID:         1,
		FQDN:               "testNode.pterodactyl.io",
		Scheme:             pteroapp.HTTPS,
		Memory:             1024,
		MemoryOverallocate: 0,
		Disk:               1000,
		DiskOverallocate:   0,
		UploadSize:         100,
		DaemonSFTP:         2022,
		DaemonListen:       8080,
	}
	node, _, err := myPteroApp.Nodes.Create(context.Background(), nodeOpts)
	if err != nil {
		log.Fatalf("error creating node: %s\n", err)
	}
	fmt.Println(node)
}
