package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
)

type TestData struct {
	Param1 string
	Param2 string
	Param3 int
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func CRUD() {
	ctx := context.Background()
	firebaseProject := os.Getenv("FIREBASE_PROJECT")
	conf := &firebase.Config{
		ProjectID: firebaseProject,
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

func handler(w http.ResponseWriter, r *http.Request) {
	CRUD()
	fmt.Fprint(w, "Hello world")
}
