package handler

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Record struct {
	ID   primitive.ObjectID
	Name string
}

//Create creating a record in a mongo
func CreateRecord(record *Record) (primitive.ObjectID, error) {
	client, ctx, cancel := DBConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	record.ID = primitive.NewObjectID()

	result, err := client.Database("airfile").Collection("records").InsertOne(ctx, record)
	if err != nil {
		log.Printf("Could not create record: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

// Gwet record from db
func GetRecord(id string) (string, error) {
	client, ctx, cancel := DBConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	objID, err := primitive.ObjectIDFromHex(id)

	result := client.Database("airfile").Collection(("records")).FindOne(ctx, bson.M{"_id": objID})

	var record Record

	err = result.Decode(&record)

	return record.Name, err
}
