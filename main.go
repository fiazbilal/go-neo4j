package main

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	ctx := context.Background()

	// Replace with your Neo4j database credentials
	dbUri := "neo4j://neo4j:7687" // or your remote URI
	dbUser := "neo4j"
	dbPassword := "password"

	// Create a new driver instance
	driver, err := neo4j.NewDriverWithContext(dbUri, neo4j.BasicAuth(dbUser, dbPassword, ""))
	if err != nil {
		log.Fatalf("Failed to create driver: %v", err)
	}
	defer driver.Close(ctx)

	// Verify connectivity
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		log.Fatalf("Failed to verify connectivity: %v", err)
	}
	fmt.Println("Connection established.")

	// Create a new session
	session := driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	// Run a Cypher query to create a new Person node
	personName := "Thor"
	_, err = session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(
			ctx,
			"CREATE (p:Person {name: $name}) RETURN p",
			map[string]any{"name": personName},
		)
		return nil, err
	})

	if err != nil {
		log.Fatalf("Failed to create person: %v", err)
	}

	fmt.Printf("Created person: %s\n", personName)
}
