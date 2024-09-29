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

	// Perform CRUD operations
	personName := "Fiaz"
	CreatePerson(session, personName)
	ReadPerson(session, personName)
	UpdatePerson(session, personName, "Fiaz Updated")
	ReadPerson(session, "Fiaz Updated")
	DeletePerson(session, "Fiaz Updated")
}

// Create a new person
func CreatePerson(session neo4j.SessionWithContext, name string) {
	_, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(
			context.Background(),
			"CREATE (p:Person {name: $name}) RETURN p",
			map[string]any{"name": name},
		)
		return nil, err
	})
	if err != nil {
		log.Fatalf("Failed to create person: %v", err)
	}
	fmt.Printf("Created person: %s\n", name)
}

// Read a person
func ReadPerson(session neo4j.SessionWithContext, name string) {
	result, err := session.ExecuteRead(context.Background(), func(tx neo4j.ManagedTransaction) (any, error) {
		record, err := tx.Run(
			context.Background(),
			"MATCH (p:Person {name: $name}) RETURN p",
			map[string]any{"name": name},
		)
		if err != nil {
			return nil, err
		}
		if !record.Next(context.Background()) {
			return nil, nil // Person not found
		}
		return record.Record().Values[0], nil
	})
	if err != nil {
		log.Fatalf("Failed to read person: %v", err)
	}
	if result == nil {
		fmt.Printf("Person not found: %s\n", name)
	} else {
		fmt.Printf("Found person: %v\n", result)
	}
}

// Update a person
func UpdatePerson(session neo4j.SessionWithContext, oldName, newName string) {
	_, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(
			context.Background(),
			"MATCH (p:Person {name: $oldName}) SET p.name = $newName RETURN p",
			map[string]any{"oldName": oldName, "newName": newName},
		)
		return nil, err
	})
	if err != nil {
		log.Fatalf("Failed to update person: %v", err)
	}
	fmt.Printf("Updated person: %s to %s\n", oldName, newName)
}

// Delete a person
func DeletePerson(session neo4j.SessionWithContext, name string) {
	_, err := session.ExecuteWrite(context.Background(), func(tx neo4j.ManagedTransaction) (any, error) {
		_, err := tx.Run(
			context.Background(),
			"MATCH (p:Person {name: $name}) DELETE p",
			map[string]any{"name": name},
		)
		return nil, err
	})
	if err != nil {
		log.Fatalf("Failed to delete person: %v", err)
	}
	fmt.Printf("Deleted person: %s\n", name)
}
