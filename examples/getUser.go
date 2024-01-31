package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/elasthicc/pterodactyl-go/pterogo/pteroapp"
)

func GetUser() {

	url := "panel_url"
	token := "panel_token"
	myPteroApp := pteroapp.NewApplication(pteroapp.WithEndpoint(url), pteroapp.WithToken(token))

	// Get user by ID
	user, _, err := myPteroApp.Users.GetByID(context.Background(), 1)
	if err != nil {
		log.Fatalf("error retirving user: %s\n", err)
	}
	fmt.Printf("user ID is: %d\n", user.Attributes.ID)

}
