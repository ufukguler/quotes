package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LoginResponseDTO struct {
	Device   string `json:"device"`
	DeviceId string `json:"deviceId"`
	Token    string `json:"token"`
	ExpireAt int64  `json:"expireAt"`
}

type QuoteResponseDTO struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        string             `json:"name"`
	Source      string             `json:"source"`
	FavCount    int                `json:"favCount"`
	IsUserLiked bool               `json:"isUserLiked"`
}

type UserQuotePageResponse struct {
	Data      []QuoteUser `json:"data"`
	Page      int64       `json:"page"`
	Size      int64       `json:"size"`
	TotalPage int64       `json:"totalPage"`
}

type CurrentUser struct {
	Id        primitive.ObjectID `json:"id"`
	DeviceId  string             `json:"deviceId"`
	CreatedAt time.Time          `json:"createdAt"`
}

type QuoteUserResponseDTO struct {
	Quote       QuoteUser `json:"quote"`
	IsUserLiked bool      `json:"isUserLiked"`
}

type QuoteUser struct {
	Id     primitive.ObjectID `bson:"_id" json:"id"`
	Name   string             `json:"name" bson:"name"`
	Source string             `json:"source" bson:"source"`
}

type Stats struct {
	UserCount     int64           `json:"total_user"`
	QuoteCount    int64           `json:"total_quote"`
	CategoryCount int64           `json:"total_category"`
	CategoryStats []CategoryStats `json:"category_stats"`
}
type CategoryStats struct {
	Name       string `json:"name"`
	QuoteCount int32  `json:"count"`
}
