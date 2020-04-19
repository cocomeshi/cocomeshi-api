package repository

import (
	"context"

	"github.com/cocomeshi/cocomeshi-api/domain/model"
)

type RestaurantRepository interface {
	FindAll(ctx context.Context) ([]*model.Restaurant, error)
}
