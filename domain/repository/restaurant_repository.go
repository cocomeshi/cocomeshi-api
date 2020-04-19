package repository

import (
	"github.com/cocomeshi/cocomeshi-api/domain/model"
	"github.com/labstack/echo"
)

type RestaurantRepository interface {
	FindAll(ctx *echo.Context) ([]model.Restaurant, error)
}
