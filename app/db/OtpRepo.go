package db

import (
	"codekar/app/models"
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func generateOTP() string {
	otp := ""
	for i := 0; i < 6; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(10))
		otp += strconv.Itoa(int(num.Int64()))
	}
	return otp
}

func SetOtp(email string) (string, error) {
	var otpObj models.OtpReq
	collection := dbClient.Database(dbName).Collection("otp")
	filter := bson.M{"email": email}
	bsonData := collection.FindOne(context.Background(), filter)
	err := bsonData.Decode(&otpObj)
	randOtp := generateOTP()
	data := bson.M{"email": email, "otp": randOtp, "createdat": time.Now().String()}
	if err != nil && err.Error() == mongo.ErrNoDocuments.Error() {
		_, err = collection.InsertOne(context.Background(), data)
		if err != nil {
			return "", err
		}
	} else if err != nil {
		return "", err
	} else {
		update := bson.M{"$set": data}
		_, err := collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "", err
		}
	}
	return randOtp, nil
}

func ValidateOtp(req models.OtpReq) error {
	var otpObj models.OtpReq
	collection := dbClient.Database(dbName).Collection("otp")
	filter := bson.M{"email": req.Email}
	bsonData := collection.FindOne(context.Background(), filter)
	err := bsonData.Decode(&otpObj)
	if err != nil && err.Error() == mongo.ErrNoDocuments.Error() {
		return errors.New("NO_OTP_GENERATION_BY_THIS_NAME")
	} else if err != nil {
		return err
	}
	fmt.Println(req)
	if otpObj.Otp == req.Otp {
		return nil
	} else {
		return errors.New("INVALID_OTP")
	}
}
