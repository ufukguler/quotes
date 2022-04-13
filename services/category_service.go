package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quotes/database"
	"quotes/models"
	"strings"
)

func FindAllCategories(categories *[]models.Category, lang string) error {
	filter := bson.M{"lang": lang}
	return database.Database.Find(ColCategory, filter, categories)
}

func FindCategoryById(id primitive.ObjectID, category *models.Category) error {
	filter := bson.M{"_id": id}
	return database.Database.FindOne(ColCategory, filter, category)
}

func FindCategoryByName(name string, category *models.Category) error {
	filter := bson.M{"name": name}
	return database.Database.FindOne(ColCategory, filter, category)
}

func FindCategoryByIdLang(id primitive.ObjectID, category *models.Category, lang string) error {
	filter := bson.M{"_id": id, "lang": lang}
	return database.Database.FindOne(ColCategory, filter, category)
}

func SaveCategory(dto models.CategoryDTO) error {
	title := strings.Title(strings.ToLower(dto.Name))
	filter := bson.M{"name": title}
	exist, err := database.Database.IsDocExist(ColCategory, filter)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("a category with given name already exists")
	}

	item := models.Category{
		Id:              primitive.NewObjectID(),
		Name:            title,
		Lang:            dto.Lang,
		BackgroundColor: dto.BackgroundColor,
		GradientColors:  dto.GradientColors,
		ImageSource:     dto.ImageSource,
	}
	return database.Database.InsertOne(ColCategory, item)
}
