package database

import (
	"context"
	"crudMongo/models"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"

func CreateConnection() (client *mongo.Client, err error) {

	clientOption := options.Client().ApplyURI(connectionString)

	client, err = mongo.Connect(context.TODO(), clientOption)
	if err != nil {

		return &mongo.Client{}, err
	}

	return client, nil
}

func CreateMovieWatchList(movie models.Netflix) (bool, error) {

	client, err := CreateConnection()
	if err != nil {
		return false, err
	}

	result, err := client.Database("netflix").Collection("watchList").InsertOne(context.Background(), movie)
	if err != nil {
		return false, err
	}

	fmt.Println("data got inserted :", result.InsertedID)

	return true, nil
}

func CreateMultipleWatchList(data []interface{}) (bool, error) {

	client, err := CreateConnection()
	if err != nil {
		return false, err
	}

	result, err := client.Database("netflix").Collection("watchList").InsertMany(context.Background(), data)
	if err != nil {
		return false, err
	}

	fmt.Printf("Data got inserted successfully: %d", result.InsertedIDs)

	return true, nil
}

func GetWatchList() (data []bson.M, err error) {

	client, err := CreateConnection()
	if err != nil {
		return []primitive.M{}, err
	}

	result, err := client.Database("netflix").Collection("watchList").Find(context.Background(), bson.M{})
	if err != nil {
		return []primitive.M{}, err
	}

	defer result.Close(context.Background())

	err = result.All(context.Background(), &data)
	if err != nil {
		return []primitive.M{}, err
	}

	return data, nil
}

func DeleteWatchList(id string) (err error) {

	client, err := CreateConnection()
	if err != nil {
		return err
	}

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := client.Database("netflix").Collection("watchList").DeleteOne(context.Background(), bson.M{"_id": objId})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("document not found")
	}

	return nil
}
