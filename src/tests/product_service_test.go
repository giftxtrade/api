package tests

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
)

func TestProductService(t *testing.T) {
	app := New(t)
	product_service := app.Service.ProductService

	t.Run("create product", func(t *testing.T) {
		t.Run("should create product", func(t *testing.T) {
			input := database.CreateProductParams{
				Title: "Product 1",
				ProductKey: "token",
				Url: "https://example.com",
				Price: "10.5",
				Rating: 4.5,
				TotalReviews: 124,
				Origin: "example",
				ImageUrl: "some-image",
				Currency: database.NullCurrencyType{
					Valid: true,
					CurrencyType: database.CurrencyTypeUSD,
				},
			}
			product, err := product_service.Querier.CreateProduct(context.Background(), input)
			if err != nil {
				t.Fatal(err)
			}

			if product.ID == 0 || product.Title != input.Title || (product.CategoryID != sql.NullInt64{}) || product.Currency != database.CurrencyTypeUSD {
				t.Fatal(product, input)
			}

			input2_category, err := product_service.Querier.CreateCategory(context.Background(), database.CreateCategoryParams{
				Name: "my category",
			})
			if err != nil {
				t.Fatal(err)
			}

			input2 := input
			input2.Title = "Product 2"
			input2.ProductKey = "token2"
			input2.Price = "1.50"
			input2.CategoryID = sql.NullInt64{
				Valid: true,
				Int64: input2_category.ID,
			}
			product2, err := product_service.Querier.CreateProduct(context.Background(), input2)
			if err != nil {
				t.Fatal(err)
			}

			if product2.Title != input2.Title || !product2.CategoryID.Valid || product2.CategoryID.Int64 != input2_category.ID {
				t.Fatal(product2, input2)
			}

			t.Run("should not create with duplicate product key", func(t *testing.T) {
				input := input
				input.Title = "Different Product"
				if _, err := product_service.Querier.CreateProduct(context.Background(), input); err == nil {
					t.Fatal("should not create product with a duplicate product_key")
				}
			})
		})
	})

	t.Run("find product", func(t *testing.T) {
		var new_product database.Product

		t.Run("find by product_key", func(t *testing.T) {
			// this should return the product created with `input`
			product, err := product_service.Querier.FindProductByProductKey(context.Background(), "token")
			if err != nil {
				t.Fatal(err, product)
			}

			if product.ProductKey != "token" || product.Title != "Product 1" {
				t.Fatal("values don't match")
			}

			input := database.CreateProductParams{
				Title: "Find Product 1",
				ProductKey: "find_product_1",
				Url: "https://example.com",
				Price: "5",
				Rating: 5,
				TotalReviews: 4,
				Currency: database.NullCurrencyType{
					Valid: true,
					CurrencyType: database.CurrencyTypeUSD,
				},
			}
			product, err = product_service.Querier.CreateProduct(context.Background(), input)
			if err != nil {
				t.Fatal(err, product)
			}

			found_product, err := product_service.Querier.FindProductByProductKey(context.Background(), product.ProductKey)
			if err != nil || !reflect.DeepEqual(found_product, product) {
				t.Fatal(err, product, found_product)
			}
			new_product = product

			if _, err := product_service.Querier.FindProductByProductKey(context.Background(), "random token"); err == nil {
				t.Fatal("product with key doesn't exist")
			}
		})

		t.Run("find by id", func(t *testing.T) {
			product, err := product_service.Querier.FindProductById(context.Background(), new_product.ID)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(product, new_product) {
				t.Fatal(product, new_product)
			}

			if _, err := product_service.Querier.FindProductById(context.Background(), 1234); err == nil {
				t.Fatal("product with key doesn't exist")
			}
		})
	})

	t.Run("should create or update", func(t *testing.T) {
		input := types.CreateProduct{
			Title: "Find Product 2",
			ProductKey: "find_product_2",
			OriginalUrl: "https://example.com",
			ImageUrl: "http://exmaple.com/image.jpg",
			Price: "5",
			Rating: 5,
			TotalReviews: 4,
			Category: "New Category",
		}
		product, created, err := product_service.UpdateOrCreate(context.Background(), input)
		if err != nil || !created {
			t.Fatal(err)
		}
		if product.ProductKey != input.ProductKey || product.Title != input.Title {
			t.Fatal("valued don't match", product, input)
		}
		if product.Origin != "example.com" {
			t.Fatal("origin is incorrect", product.Origin)
		}

		input2 := input
		input2.Title = input.Title + " (updated)"
		input2.Price = "50"
		product2, created, err := product_service.UpdateOrCreate(context.Background(), input2)
		if err != nil || created {
			t.Fatal(err)
		}

		if product2.ID != product.ID || product2.Price == product.Price {
			t.Fatal("product should be created")
		}
	})

	t.Run("filter products", func(t *testing.T) {
		t.Run("limit", func(t *testing.T) {
			products, err := product_service.Querier.FilterProducts(context.Background(), database.FilterProductsParams{
				Limit: 10,
				Search: sql.NullString{
					Valid: true,
					String: "manga",
				},
				Page: 1,
			})
			if err != nil {
				t.Fatal(err)
			}
			if len(products) != 10 {
				t.Fatal("products length is incorrect", len(products))
			}
		})
	})
}