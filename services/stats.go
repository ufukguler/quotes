package services

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quotes/database"
	"quotes/models"
)

func GetStats() models.Stats {
	return models.Stats{
		UserCount:     countCollection(ColUser),
		QuoteCount:    countCollection(ColQuote),
		CategoryCount: countCollection(ColCategory),
		CategoryStats: countQuotesByCategory(),
	}
}

func countCollection(col string) int64 {
	count, err := database.Database.CountDocuments(col, bson.M{})
	if err != nil {
		log.Errorln("Error while count collection: ", err)
		return 0
	}
	return count
}

func countQuotesByCategory() []models.CategoryStats {
	filters := make([]bson.M, 0)
	unwind := bson.M{"$unwind": "$categories"}
	group := bson.M{"$group": bson.M{
		"_id":   "$categories.name",
		"count": bson.M{"$sum": 1}},
	}

	filters = append(filters, unwind)
	filters = append(filters, group)

	var resp []primitive.D
	if err := database.Database.AggregateQuery(ColQuote, filters, &resp); err != nil {
		panic(err)
	}
	stats := make([]models.CategoryStats, 0)
	for _, d := range resp {
		stats = append(stats, models.CategoryStats{
			Name:       d.Map()["_id"].(string),
			QuoteCount: d.Map()["count"].(int32),
		})

	}
	return stats
}
