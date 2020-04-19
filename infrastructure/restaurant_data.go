package datastore

import (
	"fmt"

	"github.com/cocomeshi/cocomeshi-api/domain/model"
	"github.com/cocomeshi/cocomeshi-api/domain/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type ImplRestaurantRespository struct{}

func NewRestaurantRepository() repository.RestaurantRepository {
	return &ImplRestaurantRespository{}
}

func (r *ImplRestaurantRespository) FindAll(ctx *echo.Context) ([]model.Restaurant, error) {
	db := GetInstance()
	defer db.Close()
	rows, err := db.Query("select id, name, address, longitude, latitude, area_kind from restaurant")
	if err != nil {
		return []model.Restaurant{}, err
	}
	defer rows.Close()

	var res model.Restaurant
	var restaurantArray []model.Restaurant
	for rows.Next() {
		err := rows.Scan(&res.Id, &res.Name, &res.Address, &res.Latitude, &res.Longitude, &res.AreaKind)
		if err != nil {
			fmt.Println(err)
		}
		restaurantArray = append(restaurantArray, res)
	}

	return restaurantArray, nil
}
