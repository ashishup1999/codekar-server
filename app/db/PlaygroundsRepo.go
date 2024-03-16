package db

import (
	"codekar/app/constants"
	"codekar/app/models"
	"context"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllPgs(userName string) ([]models.Playground, error) {
	var pgs []models.Playground
	collection := dbClient.Database(dbName).Collection("playgrounds")
	filter := bson.M{"username": userName}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return []models.Playground{}, err
	}

	defer cursor.Close(context.Background())

	//iterate through cursor
	for cursor.Next(context.Background()) {
		var pg models.Playground
		err := cursor.Decode(&pg)
		if err != nil {
			return []models.Playground{}, err
		}
		pgs = append(pgs, pg)
	}
	return pgs, nil
}

func CreateNewPg(req models.CreatePgReq) (string, error) {
	newPg := models.Playground{
		Id:         uuid.New().String(),
		UserName:   req.UserName,
		PgName:     req.PgName,
		Java:       constants.JAVA_TEMPLATE,
		Cpp:        constants.CPP_TEMPLATE,
		Javascript: constants.JS_TEMPLATE,
		Go:         constants.GO_TEMPLATE,
		CreatedAt:  time.Now().String(),
		UpdatedAt:  time.Now().String(),
	}
	collection := dbClient.Database(dbName).Collection("playgrounds")
	bsonData, err := bson.Marshal(newPg)
	if err != nil {
		return "", err
	}
	result, err := collection.InsertOne(context.Background(), bsonData)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(string), nil
}

func GetPgById(pgId string) (models.Playground, error) {
	var pg models.Playground
	collection := dbClient.Database(dbName).Collection("playgrounds")
	filter := bson.M{"_id": pgId}
	bsonData := collection.FindOne(context.Background(), filter)
	err := bsonData.Decode(&pg)
	if err != nil {
		return models.Playground{}, err
	}
	return pg, nil
}

func UpdatePg(pg models.UpdatePgReq) error {
	pg.UpdatedAt = time.Now().String()
	collection := dbClient.Database(dbName).Collection("playgrounds")
	filter := bson.M{"_id": pg.PgId}
	update := bson.M{"$set": pg}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func DeletePg(Id string) error {
	collection := dbClient.Database(dbName).Collection("playgrounds")
	filter := bson.M{"_id": Id}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
