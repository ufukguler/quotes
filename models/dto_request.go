package models

type CategoryDTO struct {
	Name            string   `json:"name" bson:"name" validate:"required"`
	Lang            string   `json:"lang" bson:"lang" validate:"required"`
	BackgroundColor string   `json:"backgroundColor" bson:"backgroundColor" validate:"required"`
	GradientColors  []string `json:"gradientColors" bson:"gradientColors" validate:"required"`
	ImageSource     string   `json:"imageSource" bson:"imageSource" validate:"required"`
}

type QuoteDTO struct {
	Name        string   `json:"name" bson:"name" validate:"required"`
	Source      string   `json:"source" bson:"source" validate:"required"`
	CategoryIds []string `json:"categoryId" bson:"categoryId" validate:"required"`
	Lang        string   `json:"lang" bson:"lang" validate:"required"`
}

type RegisterDTO struct {
	DeviceId string `json:"deviceId" bson:"deviceId" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
}

type LoginDTO struct {
	Device   string `json:"device" bson:"device" validate:"required"`
	DeviceId string `json:"deviceId" bson:"deviceId" validate:"required"`
}

type AdminLoginDTO struct {
	Email    string `json:"email" bson:"email" validate:"required"`
	Password string `json:"password" bson:"password" validate:"required"`
}
