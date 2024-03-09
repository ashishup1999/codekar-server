package db

import (
	"codekar/app/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func UserExists(userName string, email string) (bool, error) {
	var userObj models.User
	collection := dbClient.Database(dbName).Collection("users")
	con1 := bson.M{"username": userName}
	con2 := bson.M{"email": email}
	filter := bson.M{"$or": []bson.M{con1, con2}}
	bsonData := collection.FindOne(context.Background(), filter)
	err := bsonData.Decode(&userObj)
	if err != nil {
		fmt.Println(err.Error())
	}
	if userObj.UserName == "" {
		return false, nil
	}
	return true, nil
}

func CreateUser(userObj models.User) error {
	collection := dbClient.Database(dbName).Collection("users")
	bsonData, err := bson.Marshal(userObj)
	if err != nil {
		return err
	}
	_, err = collection.InsertOne(context.Background(), bsonData)
	if err != nil {
		return err
	}
	return nil
}
