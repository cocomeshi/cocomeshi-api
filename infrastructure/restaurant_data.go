package datastore

import (
	"fmt"

	"github.com/cocomeshi/cocomeshi-api/domain/model"
	"github.com/cocomeshi/cocomeshi-api/domain/repository"
	"github.com/labstack/echo"
)

type ImplRestaurantRespository struct{}

func NewRestaurantRepository() repository.RestaurantRepository {
	return &ImplRestaurantRespository{}
}

func (r *ImplRestaurantRespository) FindAll(ctx echo.Context) ([]&model.Restaurant, error) {
	db := datastore.GetInstance()
	defer db.Close()
	rows, err := db.Query("select * from restaurant")
	if err != nil {
		return []string{}, err
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
