package main

import (
	"flowery-following-server/lib"
	"flowery-following-server/routes"
	"log"
)

func main() {
	client := lib.GetNeo4jClientInstance()
	err := client.Connect("neo4j://localhost:7687", "neo4j", "password") //* TODO: environmental variable is required

	if err != nil {
		panic(err) //* TODO: Better Error Handling is required.
	}

	r := routes.BootstrapRouter()
	log.Default().Println("🌟 Running Following Server in port :13456 🌟")
	err = r.Run(":13456") // TODO: Env
	if err != nil {
		panic(err) //* TODO: Better Error Handling is required.
	}

}
