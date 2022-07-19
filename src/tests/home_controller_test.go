package tests

import (
	"testing"
)

func TestHomeController(t *testing.T) {
	// req, err := http.NewRequest("GET", "/", nil)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// rr := httptest.NewRecorder()
	// home_controller := controllers.HomeController{}
	// handler := http.HandlerFunc(home_controller.Home)
	// handler.ServeHTTP(rr, req)

	// // Check the response body is what we expect.
	// expected := types.Response{Message: "GiftTrade REST API ⚡"}
	// var parsed_body types.Response
	// err = json.Unmarshal(rr.Body.Bytes(), &parsed_body)
	// if err != nil || parsed_body != expected {
	// 	t.Errorf("home handler did not return Response struct")
	// 	return
	// }
}
