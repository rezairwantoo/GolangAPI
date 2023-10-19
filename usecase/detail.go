package usecase

import (
	"case2/model"
	"case2/model/constant"
	"context"

	zlog "github.com/rs/zerolog/log"
)

func (u *Usecase) Detail(ctx context.Context, req model.DetailRequest) (model.DetailResponse, error) {
	var (
		err     error
		resp    model.DetailResponse
		product *model.Products
	)
	if product, err = u.postgreSQL.GetByID(ctx, req.ProductID); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Get products")
		resp.Message = constant.ErrMsgNotFoundDefault
		return resp, err
	}

	return model.DetailResponse{
		Message: constant.SuccessDetail,
		Data:    *product,
	}, nil
}
