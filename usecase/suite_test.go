package usecase

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	BeforeSuite(func() {
	})
	RunSpecs(t, "Usecase Suite")
}
