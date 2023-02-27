package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/memberlist"
)

func main() {
	// Create a new memberlist with some basic configuration
	config := memberlist.DefaultLocalConfig()
	config.Name = "node-1" // Set the node name
	list, err := memberlist.Create(config)
	if err != nil {
		log.Fatal("Failed to create memberlist: ", err)
	}

	// Start the gossip loop
	i := 0
	for {
		// Get a list of all the members in the cluster
		members := list.Members()

		// Print the list of members
		fmt.Println("Members:")
		for _, member := range members {
			fmt.Printf(" %s\t  %s\n", strconv.Itoa(i), member.Name)
		}
		i += 1
		// Wait for a few seconds before gossiping again
		time.Sleep(3 * time.Second)
	}
}
