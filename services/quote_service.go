package services

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"quotes/database"
	"quotes/models"
	"time"
)

func FindRandomQuote(lang string) (models.Quote, error) {
	var quotes []models.Quote
	var filter []bson.M
	filter = append(filter, bson.M{"$sample": bson.M{"size": 50}})
	filter = append(filter, bson.M{"$match": bson.M{"$and": []bson.M{{"lang": lang}}}})

	if err := findQuote(&quotes, filter); err != nil {
		return models.Quote{}, err
	}

	if len(quotes) > 0 {
		return quotes[rand.Intn(len(quotes))], nil
	}
	return models.Quote{}, errors.New("no document found")
}

func FindRandomQuoteByCategoryId(id primitive.ObjectID) (models.Quote, error) {
	var quotes []models.Quote
	var filter []bson.M
	filter = append(filter, bson.M{"$sample": bson.M{"size": 50}})
	filter = append(filter, bson.M{"$match": bson.M{"categories._id": bson.M{"$in": []primitive.ObjectID{id}}}})

	if err := findQuote(&quotes, filter); err != nil {
		return models.Quote{}, err
	}
	if len(quotes) > 0 {
		return quotes[0], nil
	}
	return models.Quote{}, errors.New("no document found")
}

func FindQuoteById(id primitive.ObjectID, quote *models.Quote) error {
	filter := bson.M{"_id": id}
	return database.Database.FindOne(ColQuote, filter, quote)
}

func SaveQuote(dto models.QuoteDTO) error {
	categories := make([]models.Category, 0)
	ids := removeDuplicateValues(dto.CategoryIds)
	for _, id := range ids {
		prim, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return errors.New("invalid language")
		}
		var category models.Category
		if err = FindCategoryByIdLang(prim, &category, dto.Lang); err != nil {
			log.Error("SaveQuote: category not found")
			return err
		}
		categories = append(categories, category)
	}

	item := models.Quote{
		Id:         primitive.NewObjectID(),
		Name:       dto.Name,
		Source:     dto.Source,
		Categories: categories,
		FavCount:   0,
		CreatedAt:  time.Now(),
		Lang:       dto.Lang,
	}
	return database.Database.InsertOne(ColQuote, item)
}

func LikeQuote(id primitive.ObjectID, deviceId string) error {
	var quote models.Quote
	if err := FindQuoteById(id, &quote); err != nil {
		return err
	}
	if err := addQuoteToLikeList(quote.Id, deviceId); err != nil {
		return err
	}

	return setNewFavCount(quote.Id, quote.FavCount+1)

}

func DisLikeQuote(id primitive.ObjectID, deviceId string) error {
	var quote models.Quote
	if err := FindQuoteById(id, &quote); err != nil {
		return err
	}
	if err := removeQuoteToLikeList(quote.Id, deviceId); err != nil {
		return err
	}
	return setNewFavCount(quote.Id, quote.FavCount-1)
}

func removeDuplicateValues(slice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func setNewFavCount(id primitive.ObjectID, newFavCount int) error {
	if newFavCount < 0 {
		return nil
	}
	filter := bson.M{"_id": id}
	updateStatement := bson.M{"$set": bson.M{"favCount": newFavCount}}
	return database.Database.UpdateOne(ColQuote, filter, updateStatement)

}

func findQuote(quotes *[]models.Quote, filter []bson.M) error {
	return database.Database.AggregateQuery(ColQuote, filter, quotes)
}
