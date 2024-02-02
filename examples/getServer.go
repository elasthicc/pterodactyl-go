package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/elasthicc/pterodactyl-go/pterogo/pteroapp"
)

func GetServerByID() {

	url := "panel_url"
	token := "panel_token"
	myPteroApp := pteroapp.NewApplication(pteroapp.WithEndpoint(url), pteroapp.WithToken(token))

	// Get server by ID
	server, _, err := myPteroApp.Servers.GetByID(context.Background(), 2)
	if err != nil {
		log.Fatalf("error retirving user: %s\n", err)
	}
	fmt.Printf("server description is: %s\n", server.Attributes.Description)

}
