package usecase

import (
	"case2/model"
	"context"
)

type ProductUsecase interface {
	Create(ctx context.Context, req model.CreateRequest) (model.CreateResponse, error)
	Detail(ctx context.Context, req model.DetailRequest) (model.DetailResponse, error)
	List(ctx context.Context, req *model.ListRequest) (*model.ListResponse, error)
	Update(ctx context.Context, req model.UpdateRequest) (model.UpdateResponse, error)
	Delete(ctx context.Context, req model.DeleteRequest) (model.DeleteResponse, error)
}
