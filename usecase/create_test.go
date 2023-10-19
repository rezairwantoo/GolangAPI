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

var _ = Describe("Create Products", func() {
	var (
		mockPostgre *mocks.MockPostgresSQLRepository
		mockUsecase ProductUsecase
		ctx         context.Context
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		ctx = context.Background()
		mockPostgre = mocks.NewMockPostgresSQLRepository(mockCtrl)
		mockUsecase = NewProductUsecase(mockPostgre)
	})

	Describe("Create Usecase", func() {
		Context("When request for Create", func() {
			It("success", func() {
				req := model.CreateRequest{
					Item:  "item 1",
					Price: 1,
				}

				mockPostgre.EXPECT().Create(ctx, req).Return(nil)
				_, err := mockUsecase.Create(ctx, req)
				Expect(err).To(BeNil())
			})
			It("return error when GetByID has failed", func() {
				req := model.CreateRequest{
					Item:  "item 1",
					Price: 1,
				}

				mockPostgre.EXPECT().Create(ctx, req).Return(errors.New("error"))
				_, err := mockUsecase.Create(ctx, req)
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
