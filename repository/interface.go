package repository

import (
	"case2/model"
	"context"
)

type PostgresSQLRepository interface {
	Create(ctx context.Context, req model.CreateRequest) error
	GetByID(ctx context.Context, productID int64) (*model.Products, error)
	List(ctx context.Context, req *model.ListRequest) ([]model.Products, error)
	CountProductBySearch(ctx context.Context, req *model.ListRequest) (int64, error)
	CountAllProduct(ctx context.Context) (int64, error)
	Update(ctx context.Context, req model.UpdateRequest) error
}
