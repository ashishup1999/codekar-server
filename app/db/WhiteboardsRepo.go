package db

import (
	"codekar/app/models"
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllWbs(userName string) ([]models.Whiteboard, error) {
	var wbs []models.Whiteboard
	collection := dbClient.Database(dbName).Collection("whiteboards")
	filter := bson.M{"username": userName}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return []models.Whiteboard{}, err
	}

	defer cursor.Close(context.Background())

	//iterate through cursor
	for cursor.Next(context.Background()) {
		var wb models.Whiteboard
		err := cursor.Decode(&wb)
		if err != nil {
			return []models.Whiteboard{}, err
		}
		wbs = append(wbs, wb)
	}
	return wbs, nil
}

func CreateNewWb(req models.CreateWbReq) (string, error) {
	newWb := models.Whiteboard{
		Id:        uuid.New().String(),
		UserName:  req.UserName,
		WbName:    req.WbName,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	collection := dbClient.Database(dbName).Collection("whiteboards")
	bsonData, err := bson.Marshal(newWb)
	if err != nil {
		return "", err
	}
	result, err := collection.InsertOne(context.Background(), bsonData)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(string), nil
}

func GetWbById(wbId string) (models.Whiteboard, error) {
	var wb models.Whiteboard
	collection := dbClient.Database(dbName).Collection("whiteboards")
	filter := bson.M{"_id": wbId}
	bsonData := collection.FindOne(context.Background(), filter)
	err := bsonData.Decode(&wb)
	if err != nil {
		return models.Whiteboard{}, err
	}
	return wb, nil
}

func UpdateWb(wb models.UpdateWbReq) error {
	wb.UpdatedAt = time.Now().String()
	collection := dbClient.Database(dbName).Collection("whiteboards")
	filter := bson.M{"_id": wb.WbId}
	update := bson.M{"$set": wb}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func DeleteWb(Id string) error {
	collection := dbClient.Database(dbName).Collection("whiteboards")
	filter := bson.M{"_id": Id}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func GetWbsByName(name string, pageNo int64) ([]models.Whiteboard, error) {
	var wbs []models.Whiteboard
	collection := dbClient.Database(dbName).Collection("whiteboards")

	//page configuration
	perPage := 10
	skip := int64((pageNo - 1) * int64(perPage))

	// Define options for pagination
	findOptions := options.Find()
	findOptions.SetLimit(int64(perPage))
	findOptions.SetSkip(skip)

	//regex building
	regex := primitive.Regex{Pattern: name, Options: "i"}

	filter := bson.M{"wbName": bson.M{"$regex": regex}}
	cursor, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		return []models.Whiteboard{}, err
	}

	defer cursor.Close(context.Background())

	//iterate through cursor
	for cursor.Next(context.Background()) {
		var wb models.Whiteboard
		err := cursor.Decode(&wb)
		if err != nil {
			return []models.Whiteboard{}, err
		}
		wbs = append(wbs, wb)
	}
	return wbs, nil
}
