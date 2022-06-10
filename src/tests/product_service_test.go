package tests

import (
	"reflect"
	"testing"

	"github.com/giftxtrade/api/src/types"
	"github.com/google/uuid"
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
			cp_input.Rating = -4
			product, err = product_service.Create(&cp_input)
			if err == nil || product != nil {
				t.Fatalf("rating cannot be negative")
			}

			cp_input = input
			cp_input.Rating = 5.1
			product, err = product_service.Create(&cp_input)
			if err == nil || product != nil {
				t.Fatalf("rating cannot be greater than 5")
			}

			cp_input = input
			cp_input.Price = 0
			product, err = product_service.Create(&cp_input)
			if err == nil || product != nil {
				t.Fatalf("price should be required")
			}

			cp_input = input
			cp_input.Price = -100.1
			product, err = product_service.Create(&cp_input)
			if err == nil || product != nil {
				t.Fatalf("price cannot be negative")
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
			input2.ProductKey = "token2"
			product2, err := product_service.Create(&input2)
			if err != nil {
				t.Fatal(err)
			}

			if product2.Title != input2.Title || !reflect.DeepEqual(product2.Category, product.Category) {
				t.Fatal(product2, input2)
			}

			t.Run("should parse url", func(t *testing.T) {
				input := input
				input.Title = "Product 3"
				input.Category = "test"
				input.OriginalUrl = "https://www.amazon.com/gp/product/B07G5MSF3G/ref=ppx_yo_dt_b_search_asin_image?ie=UTF8&psc=1"
				input.ProductKey = "x"
				product, err := product_service.Create(&input)
				if err != nil {
					t.Fatal(err)
				}

				if product.OriginalUrl != input.OriginalUrl || product.WebsiteOrigin != "www.amazon.com" {
					t.Fatal(product, input)
				}
			})

			t.Run("should not parse url", func(t *testing.T) {
				input := input
				input.Title = "Product 3"
				input.Category = "test"
				input.OriginalUrl = "invalid url"
				input.ProductKey = "y"
				if _, err := product_service.Create(&input); err == nil {
					t.Fatal("should not parse invalid url: " + input.OriginalUrl)
				}
			})

			t.Run("should not create with duplicate product key", func(t *testing.T) {
				input := input
				input.Title = "Different Product"
				if _, err := product_service.Create(&input); err == nil {
					t.Fatal("should not create product with a duplicate product_key")
				}
			})
		})
	})

	t.Run("find product", func(t *testing.T) {
		var new_product types.Product

		t.Run("find by product_key", func(t *testing.T) {
			product, err := product_service.Find("token")
			if err != nil || product == nil {
				t.Fatal(err, product)
			}

			if product.ProductKey != "token" || product.Title != "Product 1" {
				t.Fatal("values don't match")
			}

			input := types.CreateProduct{
				Title: "Find Product 1",
				ProductKey: "find_product_1",
				OriginalUrl: "https://example.com",
				Price: 5,
				Rating: 5,
				TotalReviews: 4,
				Category: "New Category",
			}
			product, create_err := product_service.Create(&input)
			if create_err != nil {
				t.Fatal(create_err)
			}
			found_product, found_err := product_service.Find(product.ProductKey)
			if found_err != nil || !reflect.DeepEqual(found_product, product) {
				t.Fatal(found_err, product, found_product)
			}
			new_product = *product

			if _, err = product_service.Find("some random token that doesn't exist"); err == nil {
				t.Fatal("product with key doesn't exist")
			}
		})

		t.Run("find by id", func(t *testing.T) {
			product, err := product_service.Find(new_product.ID.String())
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(*product, new_product) {
				t.Fatal(product, new_product)
			}

			if _, err = product_service.Find(uuid.NewString()); err == nil {
				t.Fatal("product with key doesn't exist")
			}
		})
	})

	t.Cleanup(func() {
		product_service.DB.Exec("delete from products")
	})
}