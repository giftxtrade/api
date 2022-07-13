package tests

import (
	"reflect"
	"testing"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/google/uuid"
)

func TestProductService(t *testing.T) {
	db := SetupMockProductService(t)
	product_service := services.ProductService{
		ServiceBase: services.CreateService(db, "products"),
		CategoryService: services.CategoryService{
			ServiceBase: services.CreateService(db, "categories"),
		},
	}

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
			var product types.Product

			cp_input := input
			cp_input.ProductKey = ""
			if err := product_service.Create(&cp_input, &product); err == nil {
				t.Fatalf("product_key should be required")
			}

			cp_input = input
			cp_input.Title = ""
			if err := product_service.Create(&cp_input, &product); err == nil {
				t.Fatalf("title should be required")
			}

			cp_input = input
			cp_input.OriginalUrl = ""
			if err := product_service.Create(&cp_input, &product);err == nil {
				t.Fatalf("original_url should be required")
			}

			cp_input = input
			cp_input.Rating = 0
			if err := product_service.Create(&cp_input, &product); err == nil {
				t.Fatalf("rating should be required")
			}

			cp_input = input
			cp_input.Rating = -4
			if err := product_service.Create(&cp_input, &product); err == nil {
				t.Fatalf("rating cannot be negative")
			}

			cp_input = input
			cp_input.Rating = 5.1
			if err := product_service.Create(&cp_input, &product); err == nil {
				t.Fatalf("rating cannot be greater than 5")
			}

			cp_input = input
			cp_input.Price = 0
			if err := product_service.Create(&cp_input, &product); err == nil {
				t.Fatalf("price should be required")
			}

			cp_input = input
			cp_input.Price = -100.1
			if err := product_service.Create(&cp_input, &product); err == nil {
				t.Fatalf("price cannot be negative")
			}

			cp_input = input
			cp_input.TotalReviews = 0
			if err := product_service.Create(&cp_input, &product); err == nil {
				t.Fatalf("total_reviews should be required")
			}

			cp_input = input
			cp_input.Category = ""
			if err := product_service.Create(&cp_input, &product); err == nil {
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
			var product types.Product
			if err := product_service.Create(&input, &product); err != nil {
				t.Fatal(err)
			}

			if product.Title != input.Title || product.Category.Name != input.Category {
				t.Fatal(product, input)
			}

			input2 := input
			input2.Title = "Product 2"
			input2.ProductKey = "token2"
			input2.Price = 1.50
			var product2 types.Product
			if err := product_service.Create(&input2, &product2); err != nil {
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
				input.Price = 19.99
				var product types.Product
				if err := product_service.Create(&input, &product); err != nil {
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
				var product types.Product
				if err := product_service.Create(&input, &product); err == nil {
					t.Fatal("should not parse invalid url: " + input.OriginalUrl)
				}
			})

			t.Run("should not create with duplicate product key", func(t *testing.T) {
				input := input
				input.Title = "Different Product"
				var product types.Product
				if err := product_service.Create(&input, &product); err == nil {
					t.Fatal("should not create product with a duplicate product_key")
				}
			})
		})
	})

	t.Run("find product", func(t *testing.T) {
		var new_product types.Product

		t.Run("find by product_key", func(t *testing.T) {
			var product types.Product
			if err := product_service.Find("token", &product); err != nil {
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
			
			product = types.Product{}
			if err := product_service.Create(&input, &product); err != nil {
				t.Fatal(err, product)
			}
			var found_product types.Product
			found_err := product_service.Find(input.ProductKey, &found_product)
			if found_err != nil || !reflect.DeepEqual(found_product, product) {
				t.Fatal(found_err, product, found_product)
			}
			new_product = product

			if err := product_service.Find("some random token that doesn't exist", &types.Product{}); err == nil {
				t.Fatal("product with key doesn't exist")
			}
		})

		t.Run("find by id", func(t *testing.T) {
			var product types.Product
			if err := product_service.Find(new_product.ID.String(), &product); err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(product, new_product) {
				t.Fatal(product, new_product)
			}

			if err := product_service.Find(uuid.NewString(), &types.Product{}); err == nil {
				t.Fatal("product with key doesn't exist")
			}
		})
	})

	t.Run("should create or update", func(t *testing.T) {
		input := types.CreateProduct{
			Title: "Find Product 1 (Updated)",
			ProductKey: "find_product_1",
			OriginalUrl: "https://example.com",
			Price: 5,
			Rating: 5,
			TotalReviews: 4,
			Category: "New Category",
		}
		var product types.Product
		created, err := product_service.CreateOrUpdate(&input, &product)
		if err != nil {
			t.Fatal(err)
		}
		if created {
			t.Fatal("product already exists, should not create new product")
		}
		if product.ProductKey != input.ProductKey || product.Title != input.Title {
			t.Fatal("valued don't match", product, input)
		}


		input2 := input
		input2.ProductKey = "my_new_key_input2"
		input2.Price = 50
		var product2 types.Product
		created2, err2 := product_service.CreateOrUpdate(&input2, &product2)
		if err2 != nil {
			t.Fatal(err)
		}
		if !created2 || product2.ID == product.ID {
			t.Fatal("product should be created")
		}
		if product2.ProductKey != input2.ProductKey || product2.Title != input2.Title {
			t.Fatal("valued don't match", product, input)
		}
	})

	t.Run("should filter products", func(t *testing.T) {
		filter := types.ProductFilter{
			Search: "hello",
			Limit: 1,
			Page: 1,
			MinPrice: 0,
			MaxPrice: 5000,
			Sort: "",
		}
		t.Run("filter with limit and page", func(t *testing.T) {
			filter.Limit = 1
			filter.Page = 1
			products, err := product_service.Search(filter)
			if err != nil {
				t.Fatal(err)
			}
			if len(*products) != 1 {
				t.Fatal("products array should only contain 1 element")
			}
			if (*products)[0].ProductKey != "my_new_key_input2" {
				t.Fatal("wrong first product", (*products)[0])
			}

			filter.Limit = 10
			products, err = product_service.Search(filter)
			if err != nil {
				t.Fatal(err)
			}
			if len(*products) != 5 {
				t.Fatal("total products should be 5")
			}

			filter.Limit = 5
			products2, err2 := product_service.Search(filter)
			if err2 != nil {
				t.Fatal(err2)
			}
			if len(*products2) != len(*products) {
				t.Fatal("products and products2 don't have the same length")
			}
			if !reflect.DeepEqual(*products, *products2) {
				t.Fatal("products and products2 are not equal")
			}
		})

		t.Run("filter with min and max price", func(t *testing.T) {
			filter.Limit = 5
			filter.Page = 1
			filter.MinPrice = 10000
			filter.MaxPrice = 10000
			products1, err1 := product_service.Search(filter)
			if err1 != nil {
				t.Fatal(err1)
			}
			if len(*products1) != 0 {
				t.Fatal("products1 length should be 0")
			}

			filter.MaxPrice = 20000
			products1, err1 = product_service.Search(filter)
			if err1 != nil {
				t.Fatal(err1)
			}
			if len(*products1) != 0 {
				t.Fatal("products1 length should be 0")
			}

			filter.MinPrice = 1.0
			filter.MaxPrice = 2.0
			products2, err2 := product_service.Search(filter)
			if err2 != nil {
				t.Fatal(err2)
			}
			if len(*products2) != 1 {
				t.Fatal("products2 length should be 1")
			}
			{
				product := (*products2)[0]
				var found_product types.Product
				if err := product_service.Find(product.ProductKey, &found_product); err != nil {
					t.Fatal("product with key not found", product.ProductKey)
				}
				if !reflect.DeepEqual(product, found_product) {
					t.Fatal(product, found_product)
				}
				if !(product.Price >= filter.MinPrice || product.Price <= filter.MaxPrice) {
					t.Fatal("price does not match")
				}
			}
		})
	})

	t.Cleanup(func() {
		product_service.DB.Exec("delete from products")
	})
}