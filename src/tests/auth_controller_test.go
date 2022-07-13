package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/giftxtrade/api/src/controllers"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/google/uuid"
)

func TestAuthController(t *testing.T) {
	db := SetupMockUserService(t)
	auth_controller := controllers.AuthController{
		Controller: SetupMockController(db),
	}
	user_service := auth_controller.Service.UserService
	token := auth_controller.Tokens.JwtKey

	t.Run("auth middleware", func(t *testing.T) {
		t.Run("should throw status 401", func(t *testing.T) {
			t.Run("no authorization header", func(t *testing.T) {
				req, err := http.NewRequest("GET", "/auth/profile", nil)
				if err != nil {
					t.Fatal(err)
				}
				
				rr := httptest.NewRecorder()
				handler := http.Handler(auth_controller.Controller.UseJwtAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})))
				handler.ServeHTTP(rr, req)

				if rr.Result().StatusCode != 401 {
					t.Fatal("status code must be a 401")
				}
			})

			t.Run("invalid bearer token", func(t *testing.T) {
				req, err := http.NewRequest("GET", "/auth/profile", nil)
				if err != nil {
					t.Fatal(err)
				}

				req.Header.Set("Authorization", "Bearer some-random-jwt")		
				rr := httptest.NewRecorder()
				handler := http.Handler(auth_controller.Controller.UseJwtAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})))
				handler.ServeHTTP(rr, req)

				if rr.Result().StatusCode != 401 {
					t.Fatal("status code must be a 401")
				}
			})

			t.Run("invalid jwt", func(t *testing.T) {
				req, err := http.NewRequest("GET", "/auth/profile", nil)
				if err != nil {
					t.Fatal(err)
				}

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

				req.Header.Set("Authorization", "Bearer " + jwt)
				rr := httptest.NewRecorder()
				handler := http.Handler(auth_controller.Controller.UseJwtAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(200)
				})))
				handler.ServeHTTP(rr, req)

				if rr.Result().StatusCode == 200 {
					t.Fatal("jwt claims must exist in database")
				}
			})
		})

		t.Run("should authenticate with status 200", func(t *testing.T) {
			req, err := http.NewRequest("GET", "/auth/profile", nil)
			if err != nil {
				t.Fatal(err)
			}

			var user types.User
			_, err = user_service.FindOrCreate(&types.CreateUser{
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

			req.Header.Set("Authorization", "Bearer " + jwt)
			rr := httptest.NewRecorder()
			handler := http.Handler(auth_controller.Controller.UseJwtAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
			})))
			handler.ServeHTTP(rr, req)

			if rr.Result().StatusCode == 401 {
				t.Fatal("status code must be 200 for valid JWT", jwt, rr.Result().StatusCode)
			}
		})

		t.Run("admin only authentication", func(t *testing.T) {
			t.Run("non admin user", func(t *testing.T) {
				req, err := http.NewRequest("GET", "/auth/profile", nil)
				if err != nil {
					t.Fatal(err)
				}

				var user types.User
				_, err = user_service.FindOrCreate(&types.CreateUser{
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

				req.Header.Set("Authorization", "Bearer " + jwt)
				rr := httptest.NewRecorder()
				handler := http.Handler(auth_controller.Controller.UseAdminOnly(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(200)
				})))
				handler.ServeHTTP(rr, req)

				if rr.Result().StatusCode != 401 {
					t.Fatal("user is not an admin so should not authenticate")
				}
			})

			t.Run("admin user", func(t *testing.T) {
				req, err := http.NewRequest("GET", "/auth/profile", nil)
				if err != nil {
					t.Fatal(err)
				}

				var user types.User
				_, err = user_service.FindOrCreate(&types.CreateUser{
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

				req.Header.Set("Authorization", "Bearer " + jwt)
				rr := httptest.NewRecorder()
				handler := http.Handler(auth_controller.Controller.UseAdminOnly(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(200)
				})))
				handler.ServeHTTP(rr, req)

				if rr.Result().StatusCode != 200 {
					t.Fatal("user is admin, should return status code 200.")
				}
			})
		})
	})

	t.Run("[GET] /auth/profile", func(t *testing.T) {
		t.Run("should return auth struct", func(t *testing.T) {
			req, err := http.NewRequest("GET", "/auth/profile", nil)
			if err != nil {
				t.Fatal(err)
			}

			mock_auth := types.Auth{
				User: types.User{},
				Token: token,
			}

			req = req.WithContext(context.WithValue(
				req.Context(),
				types.AuthKey,
				mock_auth,
			))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(auth_controller.GetProfile)
			handler.ServeHTTP(rr, req)

			if rr.Code != 200 {
				t.Fatal("response must be ok (200).")
			}

			var res struct {
				Data types.Auth
			}
			if json.Unmarshal(rr.Body.Bytes(), &res) != nil {
				t.Fatal("could not parse response")
			}
			if !reflect.DeepEqual(res.Data, mock_auth) {
				t.Fatal(res.Data, mock_auth)
			}
		})
	})
}