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
			input := types.CreateProduct{
				Title: "Product 1",
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
				input := types.CreateProduct{
					Title: "Product 3",
					Category: "test",
					OriginalUrl: "https://www.amazon.com/gp/product/B07G5MSF3G/ref=ppx_yo_dt_b_search_asin_image?ie=UTF8&psc=1",
				}
				product, err := product_service.Create(&input)
				if err != nil {
					t.Fatal(err)
				}

				if product.OriginalUrl != input.OriginalUrl || product.WebsiteOrigin != "www.amazon.com" {
					t.Fatal(product, input)
				}
			})
		})
	})

	t.Cleanup(func() {
		product_service.DB.Exec("delete from products")
	})
}