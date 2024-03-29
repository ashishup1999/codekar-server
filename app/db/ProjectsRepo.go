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
	newProj := models.Project{
		Id:          uuid.New().String(),
		UserName:    req.UserName,
		ProjectName: req.ProjectName,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}
	collection := dbClient.Database(dbName).Collection("projects")
	bsonData, err := bson.Marshal(newProj)
	if err != nil {
		return "", err
	}
	result, err := collection.InsertOne(context.Background(), bsonData)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(string), nil
}

func GetProjectById(projId string) (models.Project, error) {
	var project models.Project
	collection := dbClient.Database(dbName).Collection("projects")
	filter := bson.M{"_id": projId}
	bsonData := collection.FindOne(context.Background(), filter)
	err := bsonData.Decode(&project)
	if err != nil {
		return models.Project{}, err
	}
	return project, nil
}

func UpdateProject(proj models.UpdateProjReq) error {
	proj.UpdatedAt = time.Now().String()
	collection := dbClient.Database(dbName).Collection("projects")
	filter := bson.M{"_id": proj.ProjectId}
	update := bson.M{"$set": proj}
	_, err := collection.UpdateOne(context.Background(), filter, update)
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

func GetProjectsByName(name string, pageNo int64) ([]models.Project, error) {
	var projects []models.Project
	collection := dbClient.Database(dbName).Collection("projects")

	//page configuration
	perPage := 10
	skip := int64((pageNo - 1) * int64(perPage))

	// Define options for pagination
	findOptions := options.Find()
	findOptions.SetLimit(int64(perPage))
	findOptions.SetSkip(skip)

	//regex building
	regex := primitive.Regex{Pattern: name, Options: "i"}

	filter := bson.M{"projectname": bson.M{"$regex": regex}}
	cursor, err := collection.Find(context.Background(), filter, findOptions)
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
