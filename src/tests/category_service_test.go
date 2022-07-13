package tests

import (
	"reflect"
	"testing"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
)

func TestCategoryService(t *testing.T) {
	db := SetupMockCategoryService(t)
	category_service := services.CategoryService{
		ServiceBase: services.CreateService(db, "categories"),
	}

	input := types.CreateCategory{
		Name: "Fashion",
		Url: "https://example.com",
		Description: "Clothing and apparel",
	}

	t.Run("create category", func(t *testing.T) {
		t.Run("should create", func(t *testing.T) {
			t.Cleanup(func() {
				category_service.DB.Exec("delete from categories")
			})
			var input_created types.Category
			if err := category_service.Create(&input, &input_created); err != nil {
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
			var created types.Category
			err := category_service.Create(&input, &created)
			if err == nil {
				t.Fatal(err.Error())
			}
		})
	})

	t.Run("find category", func(t *testing.T) {
		t.Run("should return created category", func(t *testing.T) {
			var input_created types.Category
			if err := category_service.Create(&input, &input_created); err != nil {
				t.Fatal(err.Error())
			}

			var found_category types.Category
			if err := category_service.Find(input_created.Name, &found_category); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(found_category, input_created) {
				t.Fatal(found_category, input_created)
			}
		})
	})

	t.Cleanup(func() {
		category_service.DB.Exec("delete from categories")
	})
}