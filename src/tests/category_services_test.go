package tests

import (
	"testing"

	"github.com/giftxtrade/api/src/types"
)

func TestCategoryServices(t *testing.T) {
	category_services := SetupMockCategoryServices(t)

	t.Run("create category", func(t *testing.T) {
		t.Run("should create", func(t *testing.T) {
			input := types.CreateCategory{
				Name: "Fashion",
				Url: "https://example.com",
				Description: "Clothing and apparel",
			}
			created, err := category_services.Create(&input)
			if err != nil {
				t.Fatal(err.Error())
			}
			if created.Name != input.Name || created.Url != input.Url || created.Description != input.Description {
				t.Fatal("values should be equal", created, input)
			}
		})

		t.Run("should not create", func(t *testing.T) {
			input := types.CreateCategory{
				Name: "",
			}
			created, err := category_services.Create(&input)
			if err == nil || created != nil {
				t.Fatal(err.Error())
			}
		})
	})
}