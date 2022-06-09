package tests

import (
	"testing"

	"github.com/giftxtrade/api/src/types"
)

func TestProductService(t *testing.T) {
	product_service := SetupMockProductService(t)

	t.Run("create product", func(t *testing.T) {
		t.Run("should not create product", func(t *testing.T) {
			input := types.CreateProduct{}
			product, err := product_service.Create(&input)
			if err == nil || product != nil {
				t.Fatalf("should not create product")
			}

			input.Title = "sample product"
			product, err = product_service.Create(&input)
			if err == nil || product != nil {
				t.Fatalf(err.Error())
			}
		})

		t.Run("should create product", func(t *testing.T) {
			
		})
	})
}