package db

import (
	"codekar/app/models"
	"codekar/app/utils"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	if err != nil && err.Error() == mongo.ErrNoDocuments.Error() {
		return false, nil
	} else if err != nil {
		return false, err
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
	if err != nil && err.Error() == mongo.ErrNoDocuments.Error() {
		return false, nil
	} else if err != nil {
		return false, err
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
	if err != nil && err.Error() == mongo.ErrNoDocuments.Error() {
		return false, nil
	} else if err != nil {
		return false, err
	}
	if userObj.Email == "" {
		return false, nil
	}
	return true, nil
}

func CreateUser(userObj models.User) error {
	collection := dbClient.Database(dbName).Collection("users")
	userObj.Id = uuid.New().String()
	userObj.ConnectionReq = []string{}
	userObj.Connections = []string{}
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
		return []models.UserMata{}, err
	}

	defer cursor.Close(context.Background())

	//iterate through cursor
	for cursor.Next(context.Background()) {
		var user models.UserMata
		err := cursor.Decode(&user)
		if err != nil {
			return []models.UserMata{}, err
		}
		users = append(users, user)
	}
	return users, nil
}

func ConnectionReqs(sender string, reciever string) error {
	collection := dbClient.Database(dbName).Collection("users")

	//getting sender info
	var senderObj models.User
	senderData := collection.FindOne(context.Background(), bson.M{"username": sender})
	err := senderData.Decode(&senderObj)
	if err != nil {
		return err
	}

	//getting sender info
	var recieverObj models.User
	recieverData := collection.FindOne(context.Background(), bson.M{"username": reciever})
	err = recieverData.Decode(&recieverObj)
	if err != nil {
		return err
	}

	connStatus, err := ConnectionStatus(sender, reciever)
	if err != nil || connStatus != "NOT_CONNECTED" {
		return err
	}

	filter := bson.M{"username": reciever}
	update := bson.M{"$push": bson.M{"connReqs": senderObj.Id}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func GetAllConnectionReqs(userName string) ([]models.UserMata, string, error) {
	collection := dbClient.Database(dbName).Collection("users")

	//getting user's info
	filter := bson.M{"username": userName}
	var userObj models.User
	userData := collection.FindOne(context.Background(), filter)
	err := userData.Decode(&userObj)
	if err != nil {
		return []models.UserMata{}, "", err
	}

	//getting user's connection reqs names based on there ids in userObj.ConnectionReq
	var connReqs []models.UserMata
	filter = bson.M{"_id": bson.M{"$in": userObj.ConnectionReq}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return []models.UserMata{}, userObj.ProfileImg, err
	}
	defer cursor.Close(context.Background())

	//iterate through cursor
	for cursor.Next(context.Background()) {
		var connReqUser models.UserMata
		err := cursor.Decode(&connReqUser)
		if err != nil {
			return []models.UserMata{}, userObj.ProfileImg, err
		}
		connReqs = append(connReqs, connReqUser)
	}
	return connReqs, userObj.ProfileImg, nil
}

func ConnectionStatus(sender string, reciever string) (string, error) {
	collection := dbClient.Database(dbName).Collection("users")

	//getting sender info
	var senderObj models.User
	senderData := collection.FindOne(context.Background(), bson.M{"username": sender})
	err := senderData.Decode(&senderObj)
	if err != nil {
		return "", err
	}

	//getting sender info
	var recieverObj models.User
	recieverData := collection.FindOne(context.Background(), bson.M{"username": reciever})
	err = recieverData.Decode(&recieverObj)
	if err != nil {
		return "", err
	}

	//check if alreadyConnected
	if utils.StringIncludes(senderObj.Connections, recieverObj.Id) &&
		utils.StringIncludes(recieverObj.Connections, senderObj.Id) {
		return "CONNECTED", nil
	}

	//check if connnection request exists
	if utils.StringIncludes(senderObj.ConnectionReq, recieverObj.Id) ||
		utils.StringIncludes(recieverObj.ConnectionReq, senderObj.Id) {
		return "CONNECTION_REQUESTED", nil
	}

	//if not any of the above
	return "NOT_CONNECTED", nil
}

func RejectConnectionReqs(reciever string, sender string) error {
	collection := dbClient.Database(dbName).Collection("users")

	//getting sender info
	var senderObj models.User
	senderData := collection.FindOne(context.Background(), bson.M{"username": sender})
	err := senderData.Decode(&senderObj)
	if err != nil {
		return err
	}

	filter := bson.M{"username": reciever}
	update := bson.M{"$pull": bson.M{"connReqs": senderObj.Id}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func AddUserConnections(reciever string, sender string) error {
	collection := dbClient.Database(dbName).Collection("users")

	//getting sender info
	var senderObj models.User
	senderData := collection.FindOne(context.Background(), bson.M{"username": sender})
	err := senderData.Decode(&senderObj)
	if err != nil {
		return err
	}

	//getting sender info
	var recieverObj models.User
	recieverData := collection.FindOne(context.Background(), bson.M{"username": reciever})
	err = recieverData.Decode(&recieverObj)
	if err != nil {
		return err
	}

	//removing connection reques in reciever's info
	filterConnReq := bson.M{"username": reciever}
	updateConnReq := bson.M{"$pull": bson.M{"connReqs": senderObj.Id}}
	_, err = collection.UpdateOne(context.Background(), filterConnReq, updateConnReq)
	if err != nil {
		return err
	}

	//removing connection reques in sender's info
	filterConnReq = bson.M{"username": sender}
	updateConnReq = bson.M{"$pull": bson.M{"connReqs": recieverObj.Id}}
	_, err = collection.UpdateOne(context.Background(), filterConnReq, updateConnReq)
	if err != nil {
		return err
	}

	//update connection in reciever's info
	filter1 := bson.M{"username": reciever}
	update1 := bson.M{"$push": bson.M{"connections": senderObj.Id}}
	_, err = collection.UpdateOne(context.Background(), filter1, update1)
	if err != nil {
		return err
	}

	//update connection in senders info
	filter2 := bson.M{"username": sender}
	update2 := bson.M{"$push": bson.M{"connections": recieverObj.Id}}
	_, err = collection.UpdateOne(context.Background(), filter2, update2)
	if err != nil {
		return err
	}
	return nil
}

func RemoveUserConnection(reciever string, sender string) error {
	collection := dbClient.Database(dbName).Collection("users")

	//getting sender info
	var senderObj models.User
	senderData := collection.FindOne(context.Background(), bson.M{"username": sender})
	err := senderData.Decode(&senderObj)
	if err != nil {
		return err
	}

	//getting sender info
	var recieverObj models.User
	recieverData := collection.FindOne(context.Background(), bson.M{"username": reciever})
	err = recieverData.Decode(&recieverObj)
	if err != nil {
		return err
	}

	//remove sender from reciever conn
	filter1 := bson.M{"username": reciever}
	update1 := bson.M{"$pull": bson.M{"connections": senderObj.Id}}
	_, err = collection.UpdateOne(context.Background(), filter1, update1)
	if err != nil {
		return err
	}

	//remove reciever from sender conn
	filter2 := bson.M{"username": sender}
	update2 := bson.M{"$pull": bson.M{"connections": recieverObj.Id}}
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

	//getting user's connection names based on there ids in userObj.ConnectionReq
	var conns []string
	filter := bson.M{"_id": bson.M{"$in": userObj.Connections}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return []string{}, err
	}
	defer cursor.Close(context.Background())

	//iterate through cursor
	for cursor.Next(context.Background()) {
		var connUser models.User
		err := cursor.Decode(&connUser)
		if err != nil {
			return []string{}, err
		}
		conns = append(conns, connUser.UserName)
	}
	return conns, nil
}

func GetUserInfo(userName string) (models.UserMetaResp, error) {
	var userObj models.User
	collection := dbClient.Database(dbName).Collection("users")
	bsonData := collection.FindOne(context.Background(), bson.M{"username": userName})
	err := bsonData.Decode(&userObj)
	if err != nil {
		return models.UserMetaResp{}, err
	}

	//getting user's connection names based on there ids in userObj.ConnectionReq
	var conns []models.UserMata
	filter := bson.M{"_id": bson.M{"$in": userObj.Connections}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return models.UserMetaResp{}, err
	}
	defer cursor.Close(context.Background())

	//iterate through cursor
	for cursor.Next(context.Background()) {
		var connUser models.UserMata
		err := cursor.Decode(&connUser)
		if err != nil {
			return models.UserMetaResp{}, err
		}
		conns = append(conns, connUser)
	}

	//updating userObj
	var userInfo models.UserMetaResp
	userInfo.UserName = userObj.UserName
	userInfo.FullName = userObj.FullName
	userInfo.ProfileImg = userObj.ProfileImg
	userInfo.Connections = conns

	return userInfo, nil
}

func UpdateUserPasword(email string, newPass string) error {
	collection := dbClient.Database(dbName).Collection("users")
	filter := bson.M{"email": email}
	update := bson.M{"$set": bson.M{"password": newPass}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserDetails(req models.EditUserReq) (string, error) {
	collection := dbClient.Database(dbName).Collection("users")
	filter := bson.M{"username": req.UserName}

	//check if password correct
	isValidPass, err := ValidateUser(req.UserName, req.CurrPassword)
	if err != nil {
		return "", err
	} else if !isValidPass {
		return "INVALID_PASSWORD", nil
	}

	//update email
	if req.NewEmail != "" {
		flag, err := UserExistsByEmail(req.NewEmail)
		if err != nil {
			return "", err
		} else if flag {
			return "EMAIL_ALREADY_IN_USE", nil
		}
		update := bson.M{"$set": bson.M{"email": req.NewEmail}}
		_, err = collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "", err
		}
	}

	//update fullName
	if req.NewFullName != "" {
		update := bson.M{"$set": bson.M{"fullname": req.NewFullName}}
		_, err := collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "", err
		}
	}

	//update fullName
	if req.NewPassword != "" {
		update := bson.M{"$set": bson.M{"password": req.NewPassword}}
		_, err := collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "", err
		}
	}

	//update userName: atlast because of filter
	if req.NewUserName != "" {
		flag, err := UserExistsByUsername(req.NewUserName)
		if err != nil {
			return "", err
		} else if flag {
			return "USERNAME_ALREADY_IN_USE", nil
		}
		update := bson.M{"$set": bson.M{"username": req.NewUserName}}
		_, err = collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "", err
		}
	}
	return "UPDATED_SUCCESSFULY", nil
}

func UpdateProfilePicture(userName string, base64_str string) (string, error) {
	fmt.Println(userName)
	collection := dbClient.Database(dbName).Collection("users")
	filter := bson.M{"username": userName}
	updateOptions := options.Update().SetUpsert(true)
	update := bson.M{"$set": bson.M{"profileImg": base64_str}}
	_, err := collection.UpdateOne(context.Background(), filter, update, updateOptions)
	if err != nil {
		return "", err
	}

	return "PROFILE_IMG_UPDATED", nil
}
