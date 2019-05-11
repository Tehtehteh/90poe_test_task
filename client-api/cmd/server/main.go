package main

import (
	"context"
	"log"

	"t_task/client-api/api"
	"t_task/client-api/config"
	"t_task/client-api/ports"
	"t_task/proto"
)

func main() {
	const portsJsonPath = "ports.json"
	// Reading .env (if present) and creating config instance
	env, err := config.CreateConfig()
	if err != nil {
		log.Printf("Error creating config instance: %s", err)
		return
	}

	// Creating server instance
	log.Println("Creating server instance...")
	server := api.CreateServer(env)

	log.Printf("Reading %s filename to parse it...\n", portsJsonPath)
	p, err := ports.ReadPortsFromFile(portsJsonPath)
	if err != nil {
		log.Fatalf("Error reading ports from %s. Error: %s", portsJsonPath, err)
	}

	log.Println("Trying to insert ports one-by-one to PDS")
	var failedPorts []*proto.Port
	portsLength := len(p)
	for i := range p {
		log.Printf("Inserting port #%d of total %d", i+1, portsLength)
		port, err := server.PDServiceClient.Insert(context.Background(), &p[i])
		if err != nil {
			failedPorts = append(failedPorts, port)
			log.Printf("Error inserting port #%d. Message: %s", i+1, err)
		}
	}
	log.Printf("Successfully inserted ports into PDService. Number of failed port insertions: %d.", len(failedPorts))
	server.Serve()
}
