package usecase

import (
	"github.com/cocomeshi/cocomeshi-api/domain/model"
	datastore "github.com/cocomeshi/cocomeshi-api/infrastructure"
	"github.com/labstack/echo"
)

func GetRestaurantList(ctx *echo.Context) ([]model.Restaurant, error) {
	repos := datastore.NewRestaurantRepository()
	return repos.FindAll(ctx)
}
