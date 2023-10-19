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

var _ = Describe("Delete Products", func() {
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

	Describe("Delete Usecase", func() {
		Context("When request for delete product", func() {
			It("success", func() {
				req := model.DeleteRequest{
					ProductID: 1,
				}

				var updateRequest model.UpdateRequest
				updateRequest.IsActive = false
				updateRequest.Item = product.Item
				updateRequest.Price = product.Price
				updateRequest.ProductID = req.ProductID

				mockPostgre.EXPECT().GetByID(ctx, req.ProductID).Return(product, nil)
				mockPostgre.EXPECT().Update(ctx, updateRequest).Return(nil)
				_, err := mockUsecase.Delete(ctx, req)
				Expect(err).To(BeNil())
			})
			It("return error when GetByID has failed", func() {
				req := model.DeleteRequest{
					ProductID: 1,
				}

				var updateRequest model.UpdateRequest
				updateRequest.IsActive = false
				updateRequest.Item = product.Item
				updateRequest.Price = product.Price
				updateRequest.ProductID = req.ProductID

				mockPostgre.EXPECT().GetByID(ctx, req.ProductID).Return(nil, errors.New("error"))
				mockPostgre.EXPECT().Update(ctx, updateRequest).Return(nil)
				_, err := mockUsecase.Delete(ctx, req)
				Expect(err).To(BeNil())
			})

			It("return error when Update has failed", func() {
				req := model.DeleteRequest{
					ProductID: 1,
				}

				var updateRequest model.UpdateRequest
				updateRequest.IsActive = false
				updateRequest.Item = product.Item
				updateRequest.Price = product.Price
				updateRequest.ProductID = req.ProductID

				mockPostgre.EXPECT().GetByID(ctx, req.ProductID).Return(product, nil)
				mockPostgre.EXPECT().Update(ctx, updateRequest).Return(errors.New("error"))
				_, err := mockUsecase.Delete(ctx, req)
				Expect(err).To(BeNil())
			})
			It("return message already delete if product is active is false", func() {
				req := model.DeleteRequest{
					ProductID: 1,
				}
				product.IsActive = false

				mockPostgre.EXPECT().GetByID(ctx, req.ProductID).Return(product, nil)
				resp, err := mockUsecase.Delete(ctx, req)
				Expect(err).To(BeNil())
				Expect(resp.Message).To(Equal(constant.ErrProductAlreadyDeleted))
			})
		})
	})
})
