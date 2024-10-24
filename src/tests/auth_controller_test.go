package tests

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/giftxtrade/api/src/types"
	"github.com/gofiber/fiber/v2"
)

func TestAuthController(t *testing.T) {
	user_service := app.Service.UserService
	token := app.Tokens.JwtKey

	t.Run("auth middleware", func(t *testing.T) {
		t.Run("should throw status 401", func(t *testing.T) {
			t.Run("no authorization header", func(t *testing.T) {
				server.Get("/no_auth_header", controller.UseJwtAuth, func(c *fiber.Ctx) error {
					return nil
				})
				req := httptest.NewRequest("GET", "/no_auth_header", nil)
				res, err_res := server.Test(req)
				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 401 {
					t.Fatal("status code must be a 401", res.StatusCode)
				}
			})

			t.Run("invalid bearer token", func(t *testing.T) {
				server.Get("/invalid_bearer_token", controller.UseJwtAuth, func(c *fiber.Ctx) error {
					return nil
				})
				req := httptest.NewRequest("GET", "/invalid_bearer_token", nil)
				req.Header.Set("Authorization", "Bearer some-random-jwt")
				res, err_res := server.Test(req)
				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 401 {
					t.Fatal("status code must be a 401")
				}
			})

			t.Run("invalid jwt", func(t *testing.T) {
				jwt, err := user_service.GenerateJWT(token, &types.User{
					Name: "New User 1",
					Email: "new_user1@email.com",
					Active: true,
				})
				if err != nil {
					t.Fatal(err)
				}

				server.Get("/invalid_jwt", controller.UseJwtAuth, func(c *fiber.Ctx) error {
					return nil
				})
				req := httptest.NewRequest("GET", "/invalid_jwt", nil)
				req.Header.Set("Authorization", "Bearer " + jwt)

				res, err_res := server.Test(req)
				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 401 {
					t.Fatal("jwt claims must exist in database")
				}
			})
		})

		t.Run("should authenticate with status 200", func(t *testing.T) {
			user, _, err := user_service.FindOrCreate(context.Background(), types.CreateUser{
				Name: "Naruto Uzumaki",
				Email: "naruto_uzumaki@gmail.com",
				ImageUrl: "",
			})
			if err != nil {
				t.Fatal(err)
			}
			jwt, err := user_service.GenerateJWT(token, &user)
			if err != nil {
				t.Fatal(err)
			}

			server.Get("/valid_jwt", controller.UseJwtAuth, func(c *fiber.Ctx) error {
				return nil
			})
			req := httptest.NewRequest("GET", "/valid_jwt", nil)
			req.Header.Set("Authorization", "Bearer " + jwt)
			res, err_res := server.Test(req)

			if err_res != nil {
				t.Fatal(err_res)
			}
			if res.StatusCode == 401 {
				t.Fatal("status code must be 200 for valid JWT", jwt, res.StatusCode)
			}
		})

		t.Run("admin only authentication", func(t *testing.T) {
			t.Run("non admin user", func(t *testing.T) {
				user, _, err := user_service.FindOrCreate(context.Background(), types.CreateUser{
					Name: "Non Admin User",
					Email: "non_admin_user@gmail.com",
				})
				if err != nil {
					t.Fatal(err)
				}
				jwt, err := user_service.GenerateJWT(token, &user)
				if err != nil {
					t.Fatal(err)
				}

				server.Get("/non_admin_user", controller.UseAdminOnly, func(c *fiber.Ctx) error {
					return nil
				})
				req := httptest.NewRequest("GET", "/non_admin_user", nil)
				req.Header.Set("Authorization", "Bearer " + jwt)
				res, err_res := server.Test(req)
				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 401 {
					t.Fatal("user is not an admin so should not authenticate", res.StatusCode)
				}
			})

			t.Run("admin user", func(t *testing.T) {
				user, _, err := user_service.FindOrCreate(context.Background(), types.CreateUser{
					Name: "Admin User",
					Email: "admin_user@gmail.com",
				})
				if err != nil {
					t.Fatal(err)
				}
				// set user to admin
				_, err = user_service.Querier.SetUserAsAdmin(context.Background(), user.ID)
				if err != nil {
					t.Fatal(err)
				}

				jwt, err := user_service.GenerateJWT(token, &user)
				if err != nil {
					t.Fatal(err)
				}

				server.Get("/admin_user", controller.UseAdminOnly, func(c *fiber.Ctx) error {
					return nil
				})
				req := httptest.NewRequest("GET", "/admin_user", nil)
				req.Header.Set("Authorization", "Bearer " + jwt)
				res, err_res := server.Test(req)

				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 200 {
					t.Fatal("user is admin, should return status code 200.")
				}
			})
		})
	})

	t.Run("[GET] /auth/profile", func(t *testing.T) {
		t.Run("should return auth struct", func(t *testing.T) {
			user, _, err := user_service.FindOrCreate(context.Background(), types.CreateUser{
				Name: "Get Profile User",
				Email: "get_profile_user@gmail.com",
			})
			if err != nil {
				t.Fatal(err)
			}
			jwt, err := user_service.GenerateJWT(token, &user)
			if err != nil {
				t.Fatal(err)
			}

			mock_auth := types.Auth{
				Token: jwt,
				User: types.User{
					ID: user.ID,
					Name: user.Name,
					Email: user.Email,
					ImageUrl: user.ImageUrl,
					Active: user.Active,
					Phone: user.Phone,
					Admin: user.Admin,
				},
			}

			req := httptest.NewRequest("GET", "/auth/profile", nil)
			req.Header.Set("Authorization", "Bearer " + jwt)
			server.Get("/auth/profile", controller.UseJwtAuth, controller.GetProfile)
			res, err_res := server.Test(req)

			if err_res != nil {
				t.Fatal(err_res)
			}
			if res.StatusCode != 200 {
				t.Fatal("response must be ok (200).", res.StatusCode)
			}

			var body types.Auth
			if json.NewDecoder(res.Body).Decode(&body) != nil {
				t.Fatal("could not parse response")
			}
			if body.Token != mock_auth.Token || body.User.ID != mock_auth.User.ID {
				t.Fatal(body, mock_auth)
			}
		})
	})
}