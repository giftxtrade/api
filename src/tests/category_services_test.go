package tests

import (
	"reflect"
	"testing"

	"github.com/giftxtrade/api/src/types"
)

func TestCategoryServices(t *testing.T) {
	category_services := SetupMockCategoryServices(t)

	input := types.CreateCategory{
		Name: "Fashion",
		Url: "https://example.com",
		Description: "Clothing and apparel",
	}

	t.Run("create category", func(t *testing.T) {
		t.Run("should create", func(t *testing.T) {
			t.Cleanup(func() {
				category_services.DB.Exec("delete from categories")
			})
			input_created, err := category_services.Create(&input)
			if err != nil {
				t.Fatal(err.Error())
			}
			if input_created.Name != input.Name || input_created.Url != input.Url || input_created.Description != input.Description {
				t.Fatal("values should be equal", input_created, input)
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

	t.Run("find category", func(t *testing.T) {
		t.Run("should return created category", func(t *testing.T) {
			input_created, err := category_services.Create(&input)
			if err != nil {
				t.Fatal(err.Error())
			}

			found_category, err := category_services.Find(input_created.Name)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(found_category, input_created) {
				t.Fatal(found_category, input_created)
			}
		})
	})
}