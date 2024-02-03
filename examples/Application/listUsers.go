package examples

import (
	"context"
	"fmt"
	"log"

	"github.com/elasthicc/pterodactyl-go/pterogo/pteroapp"
)

func GetList() {

	url := "panel_url"
	token := "panel_token"
	myPteroApp := pteroapp.NewApplication(pteroapp.WithEndpoint(url), pteroapp.WithToken(token))

	// List all users
	userList, _, err := myPteroApp.Users.GetList(context.Background())
	if err != nil {
		log.Fatalf("error retirving user list: %s\n", err)
	}
	fmt.Printf("users are: %v\n", userList.Users)

}
