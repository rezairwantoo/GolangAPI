package usecase

import (
	"case2/mocks"
	"case2/model"
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Detail Products", func() {
	var (
		mockPostgre *mocks.MockPostgresSQLRepository
		mockUsecase ProductUsecase
		ctx         context.Context
		product     *model.Products
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		ctx = context.Background()
		mockPostgre = mocks.NewMockPostgresSQLRepository(mockCtrl)
		mockUsecase = NewProductUsecase(mockPostgre)
		product = &model.Products{
			ID:    1,
			Item:  "item",
			Price: 1,
		}
	})

	Describe("Get Detail Usecase", func() {
		Context("When request for get detail", func() {
			It("success", func() {
				req := model.DetailRequest{
					ProductID: 1,
				}

				mockPostgre.EXPECT().GetByID(ctx, req.ProductID).Return(product, nil)
				_, err := mockUsecase.Detail(ctx, req)
				Expect(err).To(BeNil())
			})
			It("return error when GetByID has failed", func() {
				req := model.DetailRequest{
					ProductID: 1,
				}

				mockPostgre.EXPECT().GetByID(ctx, req.ProductID).Return(nil, errors.New("error"))
				_, err := mockUsecase.Detail(ctx, req)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
