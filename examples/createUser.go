package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/elasthicc/pterodactyl-go/pterogo/pteroapp"
)

func CreateUser() {

	url := "panel_url"
	token := "panel_token"
	myPteroApp := pteroapp.NewApplication(pteroapp.WithEndpoint(url), pteroapp.WithToken(token))

	// Create a user
	userOpts := pteroapp.UserCreateOpts{
		Email:     "pterogo@example.com",
		Username:  "pterogo",
		FirstName: "ptero",
		LastName:  "go",
	}
	user, _, err := myPteroApp.Users.Create(context.Background(), userOpts)
	if err != nil {
		log.Fatalf("error retirving user: %s\n", err)
	}
	fmt.Printf("user email is: %s\n", user.Attributes.Email)
}
