package db

import (
	"codekar/app/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UserExistsByUsernameEmail(userName string, email string) (bool, error) {
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

func UserExistsByUsername(userName string) (bool, error) {
	var userObj models.User
	collection := dbClient.Database(dbName).Collection("users")
	filter := bson.M{"username": userName}
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

func UserExistsByEmail(email string) (bool, error) {
	var userObj models.User
	collection := dbClient.Database(dbName).Collection("users")
	filter := bson.M{"email": email}
	bsonData := collection.FindOne(context.Background(), filter)
	err := bsonData.Decode(&userObj)
	if err != nil {
		fmt.Println(err.Error())
	}
	if userObj.Email == "" {
		return false, nil
	}
	return true, nil
}

func CreateUser(userObj models.User) error {
	collection := dbClient.Database(dbName).Collection("users")
	userObj.Id = uuid.New().String()
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

func ValidateUser(userName string, password string) (bool, error) {
	var userObj models.User
	collection := dbClient.Database(dbName).Collection("users")
	bsonData := collection.FindOne(context.Background(), bson.M{"username": userName})
	err := bsonData.Decode(&userObj)
	if err != nil {
		return false, err
	} else if userObj.Password == password {
		return true, nil
	}
	return false, nil
}

func GetProfilesByName(name string, pageNo int64) ([]models.UserMata, error) {
	var users []models.UserMata
	collection := dbClient.Database(dbName).Collection("users")

	//page configuration
	perPage := 10
	skip := int64((pageNo - 1) * int64(perPage))

	// Define options for pagination
	findOptions := options.Find()
	findOptions.SetLimit(int64(perPage))
	findOptions.SetSkip(skip)

	//regex building
	regex := primitive.Regex{Pattern: name, Options: "i"}

	con1 := bson.M{"username": bson.M{"$regex": regex}}
	con2 := bson.M{"fullname": bson.M{"$regex": regex}}
	filter := bson.M{"$or": []bson.M{con1, con2}}
	cursor, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return []models.UserMata{}, err
	}

	defer cursor.Close(context.Background())

	//iterate through cursor
	for cursor.Next(context.Background()) {
		var user models.UserMata
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println(err.Error())
			return []models.UserMata{}, err
		}
		users = append(users, user)
	}
	return users, nil
}

func AddUserConnections(userName1 string, userName2 string) error {
	collection := dbClient.Database(dbName).Collection("users")
	filter1 := bson.M{"username": userName1}
	update1 := bson.M{"$push": bson.M{"connections": userName2}}
	_, err := collection.UpdateOne(context.Background(), filter1, update1)
	if err != nil {
		return err
	}
	filter2 := bson.M{"username": userName2}
	update2 := bson.M{"$push": bson.M{"connections": userName1}}
	_, err = collection.UpdateOne(context.Background(), filter2, update2)
	if err != nil {
		return err
	}
	return nil
}

func GetConnectionsByUser(userName string) ([]string, error) {
	var userObj models.User
	collection := dbClient.Database(dbName).Collection("users")
	bsonData := collection.FindOne(context.Background(), bson.M{"username": userName})
	err := bsonData.Decode(&userObj)
	if err != nil {
		return []string{}, err
	}
	return userObj.Connections, nil
}

func GetUserInfo(userName string) (models.UserMetaResp, error) {
	var userObj models.UserMetaResp
	collection := dbClient.Database(dbName).Collection("users")
	bsonData := collection.FindOne(context.Background(), bson.M{"username": userName})
	err := bsonData.Decode(&userObj)
	if err != nil {
		return models.UserMetaResp{}, err
	}
	return userObj, nil
}
