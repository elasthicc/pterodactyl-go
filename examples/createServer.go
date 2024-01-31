package examples

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/elasthicc/pterodactyl-go/pterogo/pteroapp"
)

func CreateServer() {

	url := "panel_url"
	token := "panel_token"
	myPteroApp := pteroapp.NewApplication(pteroapp.WithEndpoint(url), pteroapp.WithToken(token))

	// Create server
	serverOpts := pteroapp.ServerCreateOpts{
		Name:        "Ptero Server",
		User:        1,
		LocationID:  1,
		Node:        1,
		Nest:        1,
		Egg:         4,
		DockerImage: "ghcr.io/pterodactyl/yolks:java_17",
		Startup:     "java -Xms128M -XX:MaxRAMPercentage=95.0 -Dterminal.jline=false -Dterminal.ansi=true -jar {{SERVER_JARFILE}}",
		Environment: pteroapp.ServerEnvironment{
			MinecraftVersion: "latest",
			ServerJarFile:    "server.jar",
			BuildNumber:      "latest",
		},
		Limits: pteroapp.ServerLimits{
			Memory: 1024,
			Swap:   0,
			Disk:   5000,
			IO:     500,
			CPU:    100,
		},
		FeatureLimits: pteroapp.ServerFeatureLimits{
			Databases: 1,
			Backups:   3,
		},
		Allocation: pteroapp.ServerAllocation{
			Default: 2,
		},
	}

	server, resp, err := myPteroApp.Servers.Create(context.Background(), serverOpts) // yield 422 error
	if err != nil {

		bodyBytes, _ := io.ReadAll(resp.Body) // we don't really get anything here. the close statement might be messing it up tho
		fmt.Println(string(bodyBytes))

		log.Fatalf("error creating server: %s\n", err)
	}

	fmt.Printf("created server with ID %d\n\n", server.Attributes.ID)

	fmt.Println(server)

	fmt.Println(resp)

}
