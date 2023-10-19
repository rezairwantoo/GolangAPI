package usecase

import (
	"case2/model"
	"case2/model/constant"
	"context"

	zlog "github.com/rs/zerolog/log"
)

func (u *Usecase) Delete(ctx context.Context, req model.DeleteRequest) (model.DeleteResponse, error) {
	var (
		err           error
		resp          model.DeleteResponse
		product       *model.Products
		updateRequest model.UpdateRequest
	)

	if product, err = u.postgreSQL.GetByID(ctx, req.ProductID); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Get products")
		resp.Message = constant.ErrMsgNotFoundDefault
		return resp, nil
	}

	if !product.IsActive {
		resp.Message = constant.ErrProductAlreadyDeleted
		return resp, nil
	}

	updateRequest.IsActive = false
	updateRequest.Item = product.Item
	updateRequest.Price = product.Price
	updateRequest.ProductID = req.ProductID

	if err = u.postgreSQL.Update(ctx, updateRequest); err != nil {
		zlog.Info().Interface("error", err.Error()).Msg("Failed Delete products")
		resp.Message = constant.ErrDelete
		return resp, nil
	}

	return model.DeleteResponse{
		Message: constant.SuccessDelete,
		Data: model.ResponseDataCreate{
			Status: true,
		},
	}, nil
}
