package usecase

import (
	"case2/mocks"
	"case2/model"
	"case2/model/constant"
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Update Products", func() {
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
			ID:       1,
			Item:     "item",
			Price:    1,
			IsActive: true,
		}
	})

	Describe("Update Usecase", func() {
		Context("When request for update", func() {
			It("success", func() {
				req := model.UpdateRequest{
					ProductID: 1,
					Item:      "item 1",
					Price:     1,
					IsActive:  true,
				}

				mockPostgre.EXPECT().GetByID(ctx, req.ProductID).Return(product, nil)
				mockPostgre.EXPECT().Update(ctx, req).Return(nil)
				_, err := mockUsecase.Update(ctx, req)
				Expect(err).To(BeNil())
			})
			It("return error when GetByID has failed", func() {
				req := model.UpdateRequest{
					ProductID: 1,
					Item:      "item 1",
					Price:     1,
					IsActive:  true,
				}

				mockPostgre.EXPECT().GetByID(ctx, req.ProductID).Return(nil, errors.New("error"))
				mockPostgre.EXPECT().Update(ctx, req).Return(nil)
				_, err := mockUsecase.Update(ctx, req)
				Expect(err).ToNot(BeNil())
			})

			It("return error when Update has failed", func() {
				req := model.UpdateRequest{
					ProductID: 1,
					Item:      "item 1",
					Price:     1,
					IsActive:  true,
				}

				mockPostgre.EXPECT().GetByID(ctx, req.ProductID).Return(product, nil)
				mockPostgre.EXPECT().Update(ctx, req).Return(errors.New("error"))
				_, err := mockUsecase.Update(ctx, req)
				Expect(err).ToNot(BeNil())
			})
			It("return message already delete if product is active is false", func() {
				req := model.UpdateRequest{
					ProductID: 1,
					Item:      "item 1",
					Price:     1,
					IsActive:  true,
				}
				product.IsActive = false

				mockPostgre.EXPECT().GetByID(ctx, req.ProductID).Return(product, nil)
				resp, err := mockUsecase.Update(ctx, req)
				Expect(err).To(BeNil())
				Expect(resp.Message).To(Equal(constant.ErrProductAlreadyDeleted))
			})

			It("return message no data changes if there is no update on product", func() {
				req := model.UpdateRequest{
					ProductID: 1,
					Item:      "item 1",
					Price:     1,
					IsActive:  true,
				}
				product.IsActive = true
				product.Item = "item 1"

				mockPostgre.EXPECT().GetByID(ctx, req.ProductID).Return(product, nil)
				resp, err := mockUsecase.Update(ctx, req)
				Expect(err).To(BeNil())
				Expect(resp.Message).To(Equal(constant.SuccessNoDataChange))
			})
		})
	})
})
