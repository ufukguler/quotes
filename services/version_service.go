package services

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quotes/database"
	"quotes/models"
	"strings"
)

const (
	IOS     = "IOS"
	ANDROID = "ANDROID"
)

func GetVersion(versionType string) (models.Version, error) {
	var version models.Version
	ios := strings.EqualFold(versionType, IOS)
	android := strings.EqualFold(versionType, ANDROID)

	if ios || android {
		filter := bson.M{
			"key": strings.ToUpper(versionType),
		}
		err := database.Database.FindOne(ColVersion, filter, &version)
		return version, err
	}
	return version, errors.New("invalid device type")
}

func InitVersion() {
	checkVersionType(IOS, 2, 2)
	checkVersionType(ANDROID, 2, 2)
}

func checkVersionType(versionType string, current int64, minimum int64) {
	filter := bson.M{
		"key": versionType,
	}
	exist, err := database.Database.IsDocExist(ColVersion, filter)
	if err != nil {
		log.Error("[ERROR]: ", err.Error())
		return
	}
	if exist {
		return
	}
	version := models.Version{
		Id:      primitive.NewObjectID(),
		Key:     versionType,
		Current: current,
		Minimum: minimum,
	}
	if err = database.Database.InsertOne(ColVersion, version); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return
	}

}
