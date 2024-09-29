# Go and Neo4j POC

This project is a Proof of Concept (POC) demonstrating how to connect a Go application to a Neo4j database using Docker. The application is containerized and includes the necessary configurations to run seamlessly.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Usage](#usage)

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

## Getting Started

1. Clone the repository:

   ```bash
   git clone git@github.com:fiazbilal/go-neo4j.git
   cd go-neo4j

2. Build and start the application along with the Neo4j database:

   ```bash
   docker-compose up --build -d

3. Verify that the containers are running:

   ```bash
   docker ps

## Usage

   To run the Go application and connect to the Neo4j database, execute the following command:

   ```bash
   docker-compose exec app go run main.go

![image](https://github.com/user-attachments/assets/9eee48d5-fd68-4585-8054-fd5753dc2d26)

