package handlers

import (
	"errors"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"quotes/models"
	"quotes/services"
	"quotes/util"
)

func GetQuote(c echo.Context) error {
	lang := c.Get(Header).(string)
	if !util.IsLangValid(lang) {
		return e.ResponseError(errors.New("given language is not valid"), c)
	}
	quote, err := services.FindRandomQuote(lang)
	if err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}

	responseDTO, err := initQuoteBody(c, quote)
	if err != nil {
		return e.ResponseError(err, c)
	}
	log.Info("Get quote request responded. ID:", responseDTO.Id.Hex())
	return e.ResponseOk(responseDTO, c)
}

func GetQuoteById(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return e.ResponseError(err, c)
	}
	var quote models.Quote
	if err = services.FindQuoteById(id, &quote); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}

	responseDTO, err := initQuoteBody(c, quote)
	if err != nil {
		return e.ResponseError(err, c)
	}
	log.Info("Get quote by id request responded. ID:", responseDTO.Id.Hex())
	return e.ResponseOk(responseDTO, c)
}

func GetQuoteByCategoryId(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return e.ResponseError(err, c)
	}
	quote, err := services.FindRandomQuoteByCategoryId(id)
	if err != nil {
		return e.ResponseError(err, c)
	}

	responseDTO, err := initQuoteBody(c, quote)
	if err != nil {
		return e.ResponseError(err, c)
	}
	log.Info("Get quote by category id request responded. ID:", responseDTO.Id.Hex(), ". Category:", id.Hex())
	return e.ResponseOk(responseDTO, c)
}

func SaveQuote(c echo.Context) error {
	var dto models.QuoteDTO
	if err := c.Bind(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	if err := c.Validate(&dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}

	if !util.IsLangValid(dto.Lang) {
		return e.ResponseError(errors.New("given language is not valid"), c)
	}

	if err := services.SaveQuote(dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	log.Info("Save quote request responded.")
	return e.ResponseOkEmpty(c)
}

func LikeQuoteById(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return e.ResponseError(err, c)
	}
	deviceId, err := util.GetCurrentUserDeviceId(c)
	if err != nil {
		return e.ResponseError(err, c)
	}

	if err = services.LikeQuote(id, deviceId); err != nil {
		return e.ResponseError(err, c)
	}
	log.Info("Like quote by id request responded. ID:", id.Hex())
	return e.ResponseOk(true, c)
}

func DisLikeQuoteById(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return e.ResponseError(err, c)
	}
	deviceId, err := util.GetCurrentUserDeviceId(c)
	if err != nil {
		return e.ResponseError(err, c)
	}

	if err = services.DisLikeQuote(id, deviceId); err != nil {
		return e.ResponseError(err, c)
	}
	log.Info("Disike quote by id request responded. ID:", id.Hex())
	return e.ResponseOk(true, c)
}

func initQuoteBody(c echo.Context, quote models.Quote) (models.QuoteResponseDTO, error) {
	deviceId := c.Get("deviceId").(string)
	isLiked, err := services.IsUserLikedQuote(quote.Id, deviceId)
	if err != nil {
		log.Error("[ERROR]: ", err.Error())
		return models.QuoteResponseDTO{}, err
	}
	responseDTO := models.QuoteResponseDTO{
		Id:          quote.Id,
		Name:        quote.Name,
		Source:      quote.Source,
		FavCount:    quote.FavCount,
		IsUserLiked: isLiked,
	}
	return responseDTO, err
}
