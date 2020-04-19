package main

import (
	"net/http"

	usecase "github.com/cocomeshi/cocomeshi-api/application"
	"github.com/cocomeshi/cocomeshi-api/domain/model"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/restaurants", func(c echo.Context) error {
		datas, err := usecase.GetRestaurantList(&c)
		if err != nil {
			r := &model.Result{
				Success: false,
				Error: &model.Error{
					Code:    "",
					Message: "",
				},
				Data: nil,
			}
			return c.JSON(http.StatusInternalServerError, r)
		}
		r := &model.Result{
			Success: true,
			Error:   nil,
			Data:    datas,
		}
		return c.JSON(http.StatusOK, r)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
