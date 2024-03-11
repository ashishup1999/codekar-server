package db

import (
	"codekar/app/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllProjects(userName string) ([]models.Project, error) {
	var projects []models.Project
	collection := dbClient.Database(dbName).Collection("projects")
	filter := bson.M{"username": userName}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return []models.Project{}, err
	}

	defer cursor.Close(context.Background())

	//iterate through cursor
	for cursor.Next(context.Background()) {
		var proj models.Project
		err := cursor.Decode(&proj)
		if err != nil {
			return []models.Project{}, err
		}
		projects = append(projects, proj)
	}
	return projects, nil
}

func CreateNewProj(req models.CreateProjReq) (string, error) {
	var newProj models.Project
	newProj.UserName = req.UserName
	newProj.ProjectName = req.ProjectName
	collection := dbClient.Database(dbName).Collection("projects")
	bsonData, err := bson.Marshal(newProj)
	if err != nil {
		return "", err
	}
	result, err := collection.InsertOne(context.Background(), bsonData)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).String(), nil
}

func UpdateProject(proj models.Project) error {
	collection := dbClient.Database(dbName).Collection("projects")
	bsonData, err := bson.Marshal(proj)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": proj.Id}
	update := bson.M{"$set": bsonData}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func DeleteProject(Id string) error {
	collection := dbClient.Database(dbName).Collection("projects")
	filter := bson.M{"_id": Id}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
