package handlers

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
	"quotes/models"
	"quotes/services"
	"strconv"
)

func GetCurrentUser(c echo.Context) error {
	deviceId := c.Get("deviceId").(string)
	currentUser, err := services.GetUserWithoutLikeList(deviceId)
	if err != nil {
		return e.ResponseError(err, c)
	}
	log.Info("Get current user request responded. User ID is: ", currentUser.Id.Hex())
	return e.ResponseOk(currentUser, c)
}

func GetCurrentUserLikeList(c echo.Context) error {
	var opt options.FindOptions
	page, limit := pagination(c, &opt)

	var user models.User
	deviceId := c.Get("deviceId").(string)
	err := services.FindUserByDeviceId(deviceId, &user)
	if err != nil {
		return e.ResponseError(err, c)
	}

	quoteIds := make([]primitive.ObjectID, 0)
	arrMin := page * limit
	arrMax := (page + 1) * limit
	listSize := int64(len(user.LikeList))
	if listSize != 0 && listSize >= arrMin && listSize >= arrMax {
		quoteIds = user.LikeList[arrMin:arrMax]
	} else if listSize != 0 && listSize >= arrMin {
		quoteIds = user.LikeList[arrMin:]
	}
	// reverse the list
	quoteIdsReversed := make([]primitive.ObjectID, len(quoteIds))
	if len(quoteIds) != 0 {
		for i := len(quoteIds) - 1; i >= 0; i-- {
			quoteIdsReversed = append(quoteIdsReversed, quoteIds[i])
		}
	}

	quotes := make([]models.QuoteUser, 0)
	for i := range quoteIdsReversed {
		var quote models.Quote
		if err := services.FindQuoteById(quoteIdsReversed[i], &quote); err == nil {
			quotes = append(quotes, models.QuoteUser{
				Id:     quote.Id,
				Name:   quote.Name,
				Source: quote.Source,
			})
		}
	}
	response := models.UserQuotePageResponse{
		Data:      quotes,
		Page:      page,
		Size:      limit,
		TotalPage: getTotalPage(listSize, limit),
	}
	log.Info("Get like list request responded. User ID is: ", user.Id.Hex())
	return e.ResponseOk(response, c)
}

func getTotalPage(size, limit int64) int64 {
	var totalPageFloat = float64(size) / float64(limit)
	if math.Floor(totalPageFloat) < totalPageFloat {
		return int64(math.Ceil(totalPageFloat))
	}
	return int64(math.Floor(totalPageFloat))

}

func pagination(c echo.Context, FindOptions *options.FindOptions) (int64, int64) {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("size")
	var index int64 = 0
	var limit int64 = 10

	if pageParam != "" {
		parsed, err := strconv.ParseInt(pageParam, 10, 32)
		if err != nil || parsed < 0 {
			index = 0
		} else {
			index = parsed
		}
	}

	if limitParam != "" {
		parsed, err := strconv.ParseInt(limitParam, 10, 32)
		if err != nil || parsed <= 0 {
			limit = 10
		} else {
			limit = parsed
		}
	}
	FindOptions.SetSkip(index * limit)
	FindOptions.SetLimit(limit)
	return index, limit
}
