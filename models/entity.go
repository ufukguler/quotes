package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Category struct {
	Id              primitive.ObjectID `bson:"_id" json:"id"`
	Name            string             `json:"name" bson:"name"`
	Lang            string             `json:"lang" bson:"lang"`
	BackgroundColor string             `json:"backgroundColor" bson:"backgroundColor"`
	GradientColors  []string           `json:"gradientColors" bson:"gradientColors"`
	ImageSource     string             `json:"imageSource" bson:"imageSource"`
}

type Quote struct {
	Id         primitive.ObjectID `bson:"_id" json:"id"`
	Name       string             `json:"name" bson:"name"`
	Source     string             `json:"source" bson:"source"`
	Categories []Category         `json:"categories" bson:"categories"`
	FavCount   int                `json:"favCount" bson:"favCount"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	Lang       string             `json:"lang" bson:"lang"`
}

type User struct {
	Id        primitive.ObjectID   `bson:"_id" json:"id"`
	DeviceId  string               `json:"deviceId" bson:"deviceId"`
	Device    string               `json:"device" bson:"device"`
	LikeList  []primitive.ObjectID `json:"likeList" bson:"likeList"`
	CreatedAt time.Time            `json:"createdAt" bson:"createdAt"`
}
