package tests

import (
	"net/http/httptest"
	"testing"

	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func TestAuthController(t *testing.T) {
	db := SetupMockUserService(t)
	controller := SetupMockController(db)
	user_service := controller.Service.UserService
	token := controller.Tokens.JwtKey
	app := fiber.New()
	const path = "/auth/profile"

	t.Run("auth middleware", func(t *testing.T) {
		t.Run("should throw status 401", func(t *testing.T) {
			t.Run("no authorization header", func(t *testing.T) {
				app.Get("/no_auth_header", controller.UseJwtAuth, func(c *fiber.Ctx) error {
					return nil
				})
				req := httptest.NewRequest("GET", "/no_auth_header", nil)
				res, err_res := app.Test(req)
				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 401 {
					t.Fatal("status code must be a 401", res.StatusCode)
				}
			})

			t.Run("invalid bearer token", func(t *testing.T) {
				app.Get("/invalid_bearer_token", controller.UseJwtAuth, func(c *fiber.Ctx) error {
					return nil
				})
				req := httptest.NewRequest("GET", "/invalid_bearer_token", nil)
				req.Header.Set("Authorization", "Bearer some-random-jwt")
				res, err_res := app.Test(req)
				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 401 {
					t.Fatal("status code must be a 401")
				}
			})

			t.Run("invalid jwt", func(t *testing.T) {
				jwt, err := utils.GenerateJWT(token, &types.User{
					Base: types.Base{
						ID: uuid.New(),
					},
					Name: "New User 1",
					Email: "new_user1@email.com",
					IsActive: true,
				})
				if err != nil {
					t.Fatal(err)
				}

				app.Get("/invalid_jwt", controller.UseJwtAuth, func(c *fiber.Ctx) error {
					return nil
				})
				req := httptest.NewRequest("GET", "/invalid_jwt", nil)
				req.Header.Set("Authorization", "Bearer " + jwt)

				res, err_res := app.Test(req)
				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 401 {
					t.Fatal("jwt claims must exist in database")
				}
			})
		})

		t.Run("should authenticate with status 200", func(t *testing.T) {
			var user types.User
			_, err := user_service.FindOrCreate(&types.CreateUser{
				Name: "Naruto Uzumaki",
				Email: "naruto_uzumaki@gmail.com",
			}, &user)
			if err != nil {
				t.Fatal(err)
			}
			jwt, err := utils.GenerateJWT(token, &user)
			if err != nil {
				t.Fatal(err)
			}

			app.Get("/valid_jwt", controller.UseJwtAuth, func(c *fiber.Ctx) error {
				return nil
			})
			req := httptest.NewRequest("GET", "/valid_jwt", nil)
			req.Header.Set("Authorization", "Bearer " + jwt)
			res, err_res := app.Test(req)

			if err_res != nil {
				t.Fatal(err_res)
			}
			if res.StatusCode == 401 {
				t.Fatal("status code must be 200 for valid JWT", jwt, res.StatusCode)
			}
		})

		t.Run("admin only authentication", func(t *testing.T) {
			t.Run("non admin user", func(t *testing.T) {
				var user types.User
				_, err := user_service.FindOrCreate(&types.CreateUser{
					Name: "Non Admin User",
					Email: "non_admin_user@gmail.com",
				}, &user)
				if err != nil {
					t.Fatal(err)
				}
				jwt, err := utils.GenerateJWT(token, &user)
				if err != nil {
					t.Fatal(err)
				}

				app.Get("/non_admin_user", controller.UseAdminOnly, func(c *fiber.Ctx) error {
					return nil
				})
				req := httptest.NewRequest("GET", "/non_admin_user", nil)
				req.Header.Set("Authorization", "Bearer " + jwt)
				res, err_res := app.Test(req)
				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 401 {
					t.Fatal("user is not an admin so should not authenticate", res.StatusCode)
				}
			})

			t.Run("admin user", func(t *testing.T) {
				var user types.User
				_, err := user_service.FindOrCreate(&types.CreateUser{
					Name: "Admin User",
					Email: "admin_user@gmail.com",
				}, &user)
				if err != nil {
					t.Fatal(err)
				}
				// set user to admin
				user.IsAdmin = true
				if user_service.DB.Save(&user).Error != nil {
					t.Fatal("could not update user admin level")
				}

				jwt, err := utils.GenerateJWT(token, &user)
				if err != nil {
					t.Fatal(err)
				}

				app.Get("/admin_user", controller.UseAdminOnly, func(c *fiber.Ctx) error {
					return nil
				})
				req := httptest.NewRequest("GET", "/admin_user", nil)
				req.Header.Set("Authorization", "Bearer " + jwt)
				res, err_res := app.Test(req)

				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 200 {
					t.Fatal("user is admin, should return status code 200.")
				}
			})
		})
	})

	// t.Run("[GET] /auth/profile", func(t *testing.T) {
	// 	t.Run("should return auth struct", func(t *testing.T) {
	// 		req := httptest.NewRequest("GET", path, nil)

	// 		mock_auth := types.Auth{
	// 			User: types.User{},
	// 			Token: token,
	// 		}

	// 		req = req.WithContext(context.WithValue(
	// 			req.Context(),
	// 			types.AuthKey,
	// 			mock_auth,
	// 		))
	// 		rr := httptest.NewRecorder()
	// 		handler := http.HandlerFunc(controller.GetProfile)
	// 		handler.ServeHTTP(rr, req)

	// 		if rr.Code != 200 {
	// 			t.Fatal("response must be ok (200).")
	// 		}

	// 		var res struct {
	// 			Data types.Auth
	// 		}
	// 		if json.Unmarshal(rr.Body.Bytes(), &res) != nil {
	// 			t.Fatal("could not parse response")
	// 		}
	// 		if !reflect.DeepEqual(res.Data, mock_auth) {
	// 			t.Fatal(res.Data, mock_auth)
	// 		}
	// 	})
	// })
}