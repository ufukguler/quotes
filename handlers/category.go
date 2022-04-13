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

var e models.Response

const Header = "Accept-Language"

func GetCategories(c echo.Context) error {
	categories := make([]models.Category, 0)
	lang := c.Get(Header).(string)
	if !util.IsLangValid(lang) {
		return e.ResponseError(errors.New("given language is not valid"), c)
	}
	if err := services.FindAllCategories(&categories, lang); err != nil {
		return e.ResponseError(err, c)
	}
	log.Info("GetCategories request responded")
	return e.ResponseOk(categories, c)
}

func GetCategoryById(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return e.ResponseError(err, c)
	}
	var category models.Category

	if err = services.FindCategoryById(id, &category); err != nil {
		return e.ResponseError(err, c)
	}
	log.Info("GetCategoryById request responded, ID: ", category.Id.Hex())
	return e.ResponseOk(category, c)
}

func SaveCategory(c echo.Context) error {
	var dto models.CategoryDTO

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
	if err := services.SaveCategory(dto); err != nil {
		log.Error("[ERROR]: ", err.Error())
		return e.ResponseError(err, c)
	}
	log.Info("SaveCategory request responded. %s", dto)
	return e.ResponseOkEmpty(c)
}
