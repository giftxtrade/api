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
)

func TestAuthController(t *testing.T) {
	user_service := SetupMockUserService(t)
	auth_controller := controllers.AuthController{
		Controller: *SetupMockController(user_service.DB),
		UserServices: user_service,
	}

	t.Run("[GET] /auth/profile", func(t *testing.T) {
		t.Run("should return auth struct", func(t *testing.T) {
			req, err := http.NewRequest("GET", "/auth/profile", nil)
			if err != nil {
				t.Fatal(err)
			}

			mock_auth := types.Auth{
				User: types.User{},
				Token: "my-jwt-token",
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