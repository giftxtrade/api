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
)

func TestAuthController(t *testing.T) {
	user_service := SetupMockUserService(t)
	auth_controller := controllers.AuthController{
		Controller: *SetupMockController(user_service.DB),
		UserServices: user_service,
	}
	token := "my-jwt-token"

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