package tests

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/giftxtrade/api/src/types"
	"github.com/gofiber/fiber/v2"
)

func TestHomeController(t *testing.T) {
	app := New(t)
	controller := SetupMockController(app)
	server := fiber.New()

	server.Get("/", controller.Home)
	req := httptest.NewRequest("GET", "/", nil)
	res, err_res := server.Test(req)
	if err_res != nil {
		t.Fatal(err_res.Error())
	}

	// Check the response body is what we expect.
	expected := types.Response{Message: "GiftTrade REST API ⚡"}
	if res.StatusCode != fiber.StatusOK {
		t.Fatal("incorrect response type", res.StatusCode)
	}

	var body types.Response
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body != expected {
		t.Fatal(expected)
	}
}
