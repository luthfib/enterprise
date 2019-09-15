package main

import (
	"context"
	"enterprise/app"
	"enterprise/routes"
	"log"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {

	client, err := mongo.NewClient("mongodb://root:example@0.0.0.0:27017")
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	routersInit := routes.InitRouter(client)

	app.Init()
	// Start and run the server
	routersInit.Run(":8080")
}
