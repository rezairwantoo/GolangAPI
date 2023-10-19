package usecase

import (
	"case2/model"
	"context"
	"fmt"

	"case2/model/constant"

	"golang.org/x/sync/errgroup"
)

// List ...
func (u *Usecase) List(ctx context.Context, req *model.ListRequest) (*model.ListResponse, error) {
	var (
		resp             model.ListResponse
		err              error
		products         []model.Products
		totalGetRemisier int64
		totalAllData     int64
		eg               errgroup.Group
	)

	u.prepareRequestList(req)
	eg.Go(func() error {
		var errGetList error
		if products, errGetList = u.postgreSQL.List(ctx, req); errGetList != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		var errCountGetRemisier error
		// for count by searching value
		if totalGetRemisier, errCountGetRemisier = u.postgreSQL.CountProductBySearch(ctx, req); errCountGetRemisier != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		var errCountAllRemisier error
		// for count all products data
		if totalAllData, errCountAllRemisier = u.postgreSQL.CountAllProduct(ctx); errCountAllRemisier != nil {
			return err
		}
		return nil
	})

	errGroup := eg.Wait()
	if errGroup != nil {
		return &resp, errGroup
	}

	resp.Data.Data = products
	resp.Data.Pagination = model.Pagination{
		Total:        totalGetRemisier,
		Page:         int64(req.Page),
		TotalAllData: totalAllData,
	}
	return &resp, nil
}

func (u *Usecase) prepareRequestList(req *model.ListRequest) {
	if req.Limit <= 0 {
		req.Limit = constant.DefaultLimit
	}

	if req.Page <= 0 {
		req.Page = constant.DefaultPage
	}

	if req.Search != "" {
		req.Search = fmt.Sprintf("%%%s%%", req.Search)
	}
	req.Offset = req.Limit * (req.Page - 1)
}
