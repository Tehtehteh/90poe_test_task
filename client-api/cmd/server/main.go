package main

import (
	"context"
	"log"

	"t_task/client-api/api"
	"t_task/client-api/config"
	"t_task/client-api/parser"
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
	_, err = parser.ReadPortsFromFile(portsJsonPath, func(p *proto.Port) error {
		log.Printf("Trying to send port by %s ID over gRPC", p.ID)
		_, err := server.Client.Insert(context.Background(), p)
		if err != nil {
			return err
		}
		log.Printf("Successfully sent port over gRPC (port ID: %s)", p.ID)
		return err
	})
	if err != nil {
		log.Fatalf("Error reading parser from %s. Error: %s", portsJsonPath, err)
	}

	server.Serve()
}
