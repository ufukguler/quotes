package services

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quotes/database"
	"quotes/models"
	"quotes/util"
	"time"
)

func GenerateJWT(dto models.LoginDTO) (models.LoginResponseDTO, error) {
	var user models.User
	if err := getUser(&user, dto); err != nil {
		return models.LoginResponseDTO{}, err
	}

	tokenString, expire, err := util.GenerateJwt(user.DeviceId)
	responseDTO := models.LoginResponseDTO{
		DeviceId: user.DeviceId,
		Token:    tokenString,
		Device:   user.Device,
		ExpireAt: expire,
	}
	return responseDTO, err
}

func getUser(user *models.User, dto models.LoginDTO) error {
	filter := bson.M{"deviceId": dto.DeviceId}

	exist, err := database.Database.IsDocExist(ColUser, filter)
	if err != nil {
		return err
	}
	if exist {
		log.Info("User already exists on DB. DeviceId: ", dto.DeviceId)
		return database.Database.FindOne(ColUser, filter, user)
	}
	log.Info("Creating user on DB. DeviceId: ", dto.DeviceId)

	item := models.User{
		Id:        primitive.NewObjectID(),
		DeviceId:  dto.DeviceId,
		Device:    dto.Device,
		CreatedAt: time.Now(),
	}
	err = database.Database.InsertOne(ColUser, item)
	if err != nil {
		return err
	}

	return database.Database.FindOne(ColUser, filter, user)
}

func FindUserByDeviceId(deviceId string, user *models.User) error {
	filter := bson.M{"deviceId": deviceId}
	return database.Database.FindOne(ColUser, filter, user)
}

func addQuoteToLikeList(id primitive.ObjectID, deviceId string) error {
	var user models.User
	err := FindUserByDeviceId(deviceId, &user)
	if err != nil {
		return err
	}

	list := user.LikeList
	if contains(user.LikeList, id) == false {
		list = append(list, id)
	} else {
		return nil
	}

	updateFilter := bson.M{"deviceId": deviceId}
	updateStatement := bson.M{"$set": bson.M{"likeList": list}}
	return database.Database.UpdateOne(ColUser, updateFilter, updateStatement)
}

func removeQuoteToLikeList(id primitive.ObjectID, deviceId string) error {
	var user models.User
	err := FindUserByDeviceId(deviceId, &user)
	if err != nil {
		return err
	}

	list := user.LikeList
	if contains(user.LikeList, id) {
		for i, v := range list {
			if v == id {
				list = remove(list, i)
				break
			}
		}
	} else {
		return nil
	}

	updateFilter := bson.M{"deviceId": deviceId}
	updateStatement := bson.M{"$set": bson.M{"likeList": list}}
	return database.Database.UpdateOne(ColUser, updateFilter, updateStatement)
}

func contains(arr []primitive.ObjectID, value primitive.ObjectID) bool {
	for _, a := range arr {
		if a == value {
			return true
		}
	}
	return false
}

func remove(s []primitive.ObjectID, i int) []primitive.ObjectID {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func IsUserLikedQuote(id primitive.ObjectID, deviceId string) (bool, error) {
	filter := bson.M{
		"deviceId": deviceId,
		"likeList": bson.M{"$in": []primitive.ObjectID{id}},
	}
	var user models.User
	if err := database.Database.FindOne(ColUser, filter, &user); err != nil {
		return false, nil
	}
	return true, nil
}

func GetUserWithoutLikeList(deviceId string) (models.CurrentUser, error) {
	var user models.User
	filter := bson.M{"deviceId": deviceId}
	if err := database.Database.FindOne(ColUser, filter, &user); err != nil {
		return models.CurrentUser{}, err
	}
	currentUser := models.CurrentUser{
		Id:        user.Id,
		DeviceId:  user.DeviceId,
		CreatedAt: user.CreatedAt,
	}
	return currentUser, nil
}
