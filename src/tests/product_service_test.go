package tests

import (
	"reflect"
	"testing"

	"github.com/giftxtrade/api/src/types"
)

func TestProductService(t *testing.T) {
	product_service := SetupMockProductService(t)

	t.Run("create product", func(t *testing.T) {
		t.Run("should not create product", func(t *testing.T) {
			input := types.CreateProduct{
				Title: "p1",
				ProductKey: "token",
				OriginalUrl: "https://example.com",
				Price: 10.5,
				Rating: 4.5,
				TotalReviews: 124,
				Category: "test category 1",
			}

			cp_input := input
			cp_input.ProductKey = ""
			product, err := product_service.Create(&cp_input)
			if err == nil || product != nil {
				t.Fatalf("product_key should be required")
			}

			cp_input = input
			cp_input.Title = ""
			product, err = product_service.Create(&cp_input)
			if err == nil || product != nil {
				t.Fatalf("title should be required")
			}

			cp_input = input
			cp_input.OriginalUrl = ""
			product, err = product_service.Create(&cp_input)
			if err == nil || product != nil {
				t.Fatalf("original_url should be required")
			}

			cp_input = input
			cp_input.Rating = 0
			product, err = product_service.Create(&cp_input)
			if err == nil || product != nil {
				t.Fatalf("rating should be required")
			}

			cp_input = input
			cp_input.Price = 0
			product, err = product_service.Create(&cp_input)
			if err == nil || product != nil {
				t.Fatalf("price should be required")
			}

			cp_input = input
			cp_input.TotalReviews = 0
			product, err = product_service.Create(&cp_input)
			if err == nil || product != nil {
				t.Fatalf("total_reviews should be required")
			}

			cp_input = input
			cp_input.Category = ""
			product, err = product_service.Create(&cp_input)
			if err == nil || product != nil {
				t.Fatalf("category should be required")
			}
		})

		t.Run("should create product", func(t *testing.T) {
			input := types.CreateProduct{
				Title: "Product 1",
				ProductKey: "token",
				OriginalUrl: "https://example.com",
				Price: 10.5,
				Rating: 4.5,
				TotalReviews: 124,
				Category: "any",
			}
			product, err := product_service.Create(&input)
			if err != nil {
				t.Fatal(err)
			}

			if product.Title != input.Title || product.Category.Name != input.Category {
				t.Fatal(product, input)
			}

			input2 := input
			input2.Title = "Product 2"
			product2, err := product_service.Create(&input2)
			if err != nil {
				t.Fatal(err)
			}

			if product2.Title != input2.Title || !reflect.DeepEqual(product2.Category, product.Category) {
				t.Fatal(product2, input2)
			}

			t.Run("should parse url", func(t *testing.T) {
				input.Title = "Product 3"
				input.Category = "test"
				input.OriginalUrl = "https://www.amazon.com/gp/product/B07G5MSF3G/ref=ppx_yo_dt_b_search_asin_image?ie=UTF8&psc=1"
				product, err := product_service.Create(&input)
				if err != nil {
					t.Fatal(err)
				}

				if product.OriginalUrl != input.OriginalUrl || product.WebsiteOrigin != "www.amazon.com" {
					t.Fatal(product, input)
				}
			})

			t.Run("should not parse url", func(t *testing.T) {
				input.Title = "Product 3"
				input.Category = "test"
				input.OriginalUrl = "invalid url"
				if _, err := product_service.Create(&input); err == nil {
					t.Fatal("should not parse invalid url: " + input.OriginalUrl)
				}
			})
		})
	})

	t.Cleanup(func() {
		product_service.DB.Exec("delete from products")
	})
}