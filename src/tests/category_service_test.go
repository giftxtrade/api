package tests

import (
	"context"
	"database/sql"
	"testing"

	"github.com/giftxtrade/api/src/database"
)

func TestCategoryService(t *testing.T) {
	app := New(t)
	querier := app.Querier

	input1 := database.CreateCategoryParams{
		Name: "hello",
	}

	input2 := database.CreateCategoryParams{
		Name: "person",
		Description: sql.NullString{
			Valid: true,
			String: "some description",
		},
		CategoryUrl: sql.NullString{
			Valid: true,
			String: "https://example.com",
		},
	}

	t.Run("create category", func(t *testing.T) {
		category, err := querier.CreateCategory(context.Background(), input1)
		if err != nil && category.Name != input1.Name {
			t.Fatal(err)
		}
	})

	t.Run("find category", func(t *testing.T) {
		category, err := querier.FindCategoryByName(context.Background(), input1.Name)
		if err != nil && category.Name != input1.Name {
			t.Fatal(err)
		}
	})

	t.Run("find or create", func(t *testing.T) {
		category, err := app.Service.ProductService.FindOrCreateCategory(context.Background(), input1)
		if err != nil && category.Name != input1.Name {
			t.Fatal(err)
		}

		category2, err := app.Service.ProductService.FindOrCreateCategory(context.Background(), input2)
		if err != nil && category2.Name != input2.Name {
			t.Fatal(err)
		}
	})
}
