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

	/*
		Ports available on the node (this is just a reference)
		ID: 1| Port:25565
		ID: 2| Port:25566
		ID: 3| Port:25567
		ID: 4| Port:25568
		ID: 5| Port:25569
	*/

	// This is optional, but if you want to you can assign other ports as well
	addtlAllocation := []int{2, 3, 4, 5}

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
			Default:               1,
			AdditionalAllocations: addtlAllocation,
		},
		StartOnCompletion: true,
	}

	server, resp, err := myPteroApp.Servers.Create(context.Background(), serverOpts)
	if err != nil {

		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Println(string(bodyBytes))

		log.Fatalf("error creating server: %s\n", err)
	}

	fmt.Printf("created server with ID %d\n\n", server.Attributes.ID)

	fmt.Println(server)

	fmt.Println(resp)

}
