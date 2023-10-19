package usecase

import (
	"case2/model"
	"case2/model/constant"
	"context"

	zlog "github.com/rs/zerolog/log"
)

func (u *Usecase) Update(ctx context.Context, req model.UpdateRequest) (model.UpdateResponse, error) {
	var (
		err                error
		resp               model.UpdateResponse
		product            *model.Products
		noDataItemChanges  bool
		noDataPriceChanges bool
	)

	if product, err = u.postgreSQL.GetByID(ctx, req.ProductID); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Get products")
		resp.Message = constant.ErrMsgNotFoundDefault
		return resp, err
	}

	if !product.IsActive {
		resp.Message = constant.ErrProductAlreadyDeleted
		return resp, nil
	}

	if product.Item == req.Item {
		noDataItemChanges = true
	}

	if product.Price == req.Price {
		noDataPriceChanges = true
	}

	if noDataItemChanges && noDataPriceChanges {
		resp.Message = constant.SuccessNoDataChange
		return resp, nil
	}

	if err = u.postgreSQL.Update(ctx, req); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Update products")
		resp.Message = constant.ErrUpdate
		return resp, err
	}

	return model.UpdateResponse{
		Message: constant.SuccessDetail,
		Data: model.ResponseDataCreate{
			Status: true,
		},
	}, nil
}
