package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
)

type TestData struct {
	Param1 string
	Param2 string
	Param3 int
}

func main() {
	ctx := context.Background()
	conf := &firebase.Config{
		ProjectID: "test-firestore-ca8e7",
	}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalf("error initializing app:\n %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting client:\n %v", err)
	}

	created_data := TestData{
		Param1: "val1",
		Param2: "val2",
		Param3: 3,
	}

	log.Println("Creating data")
	ref, res, err := client.Collection("todo").Add(ctx, created_data)

	if err != nil {
		log.Printf("error creating data: \n %#v", err)
	}
	log.Printf("Reference: %+v", ref)
	log.Printf("Result: %+v", res)

	log.Println("Getting data")
	doc, err := client.Collection("todo").Doc(ref.ID).Get(ctx)

	res_data := TestData{}
	if err != nil {
		log.Printf("error getting data:\n%v", err)
	}
	if err = doc.DataTo(&res_data); err != nil {
		log.Printf("Error parsing data: %v", err)
	}
	log.Printf("Result Read: %#v", res_data)

	log.Println("Updating data")
	updated_data := created_data
	updated_data.Param1 = "New val1"

	updated_res, err := client.Collection("todo").Doc(ref.ID).Set(ctx, updated_data)

	if err != nil {
		log.Printf("Error updating data:\n%#v", err)
	}

	log.Printf("Updated Data: %#v", updated_res)

	log.Println("Deleting data")

	delete_this_data := TestData{
		Param1: "Delete this data",
	}

	deleted_data_ref, deleted_data_res, err := client.Collection("todo").Add(ctx, delete_this_data)
	log.Printf("Created data to delete: %v", deleted_data_res)
	if err != nil {
		log.Printf("Error creating data to delete, %v", err)
	}

	delete_res, err := deleted_data_ref.Delete(ctx)
	if err != nil {
		log.Printf("Error deleting data, %v", err)
	}
	log.Printf("Deleted data: %v", delete_res)
}
