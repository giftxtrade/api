package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/giftxtrade/api/src/controllers"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func create_participants(n int) []types.CreateParticipant {
	participants := make([]types.CreateParticipant, n)
	for i := 0; i < n; i++ {
		id := i + 1
		participants[i] = types.CreateParticipant{
			Name: fmt.Sprintf("Participant #%d", id),
			Email: fmt.Sprintf("participant_%d@example.com", id),
			Participates: true,
		}
	}
	return participants
}

func TestEventService(t *testing.T) {
	user_service := app.Service.UserService
	event_service := app.Service.EventService
	user_1, _, err := app.Service.UserService.FindOrCreate(context.Background(), types.CreateUser{
		Name: "Test User",
		Email: "testuser@email.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	user_2, _, err := app.Service.UserService.FindOrCreate(context.Background(), types.CreateUser{
		Name: "Some other random user",
		Email: "thisrandomuser@email.com",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Run("create event", func(t *testing.T) {
		t.Run("correct params", func(t *testing.T) {
			input := types.CreateEvent{
				Name: "Event 1",
				Budget: 100.00,
				DrawAt: time.Now(),
				CloseAt: time.Now().Add(time.Hour * 24 * 30),
				Participants: append(create_participants(5), types.CreateParticipant{
					Name: user_1.Name,
					Email: user_1.Email,
					Organizer: true,
					Participates: false,
				}),
			}
			new_event, err := event_service.CreateEvent(context.Background(), &user_1, input)
			if err != nil {
				t.Fatal(err)
			}

			if len(new_event.Participants) != 6 {
				t.Fatal("not all participants were inserted")
			}
			if new_event.Budget != fmt.Sprintf("$%.2f", input.Budget) {
				t.Fatalf("values don't match %s %.2f", new_event.Budget, input.Budget)
			}

			var mp types.Participant
			for _, p := range new_event.Participants {
				if p.Email != user_1.Email {
					continue
				}
				mp = p
			}
			if mp.Accepted != true && mp.UserID != user_1.ID {
				t.Fatalf("main participant accepted or user_id fields are incorrect %#v", mp)
			}
		})

		t.Run("main participant not marked as organizer", func(t *testing.T) {
			event := types.CreateEvent{
				Name: "Event 1",
				Budget: 100.00,
				DrawAt: time.Now(),
				CloseAt: time.Now().Add(time.Hour * 24 * 30),
				Participants: append(create_participants(5), types.CreateParticipant{
					Name: user_1.Name,
					Email: user_1.Email,
					Organizer: false,
					Participates: true,
				}),
			}
			_, err := event_service.CreateEvent(context.Background(), &user_1, event)
			if err == nil {
				t.Fatal("event should not be created. main participant is not marked 'organizer'")
			}

			event.Participants = create_participants(20)
			_, err = event_service.CreateEvent(context.Background(), &user_1, event)
			if err == nil {
				t.Fatal("event should not be created. main participant was not provided")
			}
		})
	})

	t.Run("event authentication", func(t *testing.T) {
		token := app.Tokens.JwtKey
		user_1_jwt, _ := user_service.GenerateJWT(token, &user_1)
		user_2_jwt, _ := user_service.GenerateJWT(token, &user_2)

		user2_event, err := event_service.CreateEvent(context.Background(), &user_2, types.CreateEvent{
			Name: "UseEventAuthWithParam Test Event",
			Budget: 200,
			DrawAt: time.Now(),
			CloseAt: time.Now().Add(time.Hour * 24 * 30), // 30 days from now
			Participants: append(create_participants(20), types.CreateParticipant{
				Email: user_2.Email,
				Name: user_2.Name,
				Organizer: true,
			}),
		})
		if err != nil {
			t.Fatal(err)
		}

		participant := user2_event.Participants[1]
		user_3, _, _ := user_service.FindOrCreate(context.Background(), types.CreateUser{
			Name: participant.Name,
			Email: participant.Email,
		})
		user_3_jwt, _ := user_service.GenerateJWT(token, &user_3)

		t.Run("UseEventAuthWithParam", func(t *testing.T) {
			t.Run("non numeric event_id", func(t *testing.T) {
				req := httptest.NewRequest("GET", "/event/abc123", nil)
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user_1_jwt))
				server.Get("/event/:event_id", controller.UseJwtAuth, controller.UseEventAuthWithParam, func(c *fiber.Ctx) error {
					return nil
				})				
				res, err_res := server.Test(req)
				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 400 {
					t.Fatal("status code must be a 400", res.StatusCode)
				}
			})

			t.Run("incorrect numeric event_id", func(t *testing.T) {
				req := httptest.NewRequest("GET", "/event/235421", nil)
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user_1_jwt))
				server.Get("/event/:event_id", controller.UseJwtAuth, controller.UseEventAuthWithParam, func(c *fiber.Ctx) error {
					return nil
				})				
				res, err_res := server.Test(req)
				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 404 {
					t.Fatal("status code must be a 400", res.StatusCode)
				}
			})

			t.Run("correct event_id", func(t *testing.T) {
				// test with user_1's auth info
				// since this user shouldn't be on the participants list
				req := httptest.NewRequest("GET", "/event/" + fmt.Sprintf("%d", user2_event.ID), nil)
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user_1_jwt))
				server.Get("/event/:event_id", controller.UseJwtAuth, controller.UseEventAuthWithParam, func(c *fiber.Ctx) error {
					return nil
				})
				res, err_res := server.Test(req)
				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 404 {
					t.Fatal("status code must be a 400", res.StatusCode)
				}

				// test with valid user_2's auth info
				req = httptest.NewRequest("GET", "/event/" + fmt.Sprintf("%d", user2_event.ID) + "/someotherthing", nil)
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user_2_jwt))
				server.Get("/event/:event_id/someotherthing", controller.UseJwtAuth, controller.UseEventAuthWithParam, func(c *fiber.Ctx) error {
					event_id := c.UserContext().Value(controllers.EVENT_ID_PARAM_KEY).(int64)
					return utils.DataResponse(c, map[string]int64{"event_id": event_id})
				})				
				res, err_res = server.Test(req)
				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 200 {
					t.Fatal("status code must be a 200", res.StatusCode)
				}
				var body map[string]int64
				if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
					t.Fatal("could not parse body", err.Error())
				}
				if body["event_id"] != user2_event.ID {
					t.Fatal("event id is incorrect", body["event_id"], user2_event.ID)
				}

				// test with unaccepted invite participant user
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user_3_jwt))
				res, err_res = server.Test(req)
				if err_res != nil {
					t.Fatal(err_res)
				}
				if res.StatusCode != 200 {
					t.Fatal("status code must be a 200", res.StatusCode)
				}
				if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
					t.Fatal("could not parse body", err.Error())
				}
				if body["event_id"] != user2_event.ID {
					t.Fatal("event id is incorrect", body["event_id"], user2_event.ID)
				}
			})
		})
	
		t.Run("UseEventOrganizerAuthWithParam", func(t *testing.T) {
			// test with user_3's auth with no organizer permissions
			req := httptest.NewRequest("GET", "/event/" + fmt.Sprint(user2_event.ID) + "/another-route", nil)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user_3_jwt))
			server.Get("/event/:event_id/another-route", controller.UseJwtAuth, controller.UseEventOrganizerAuthWithParam, func(c *fiber.Ctx) error {
				event_id := c.UserContext().Value(controllers.EVENT_ID_PARAM_KEY).(int64)
				return utils.DataResponse(c, map[string]int64{"event_id": event_id})
			})
			res, err_res := server.Test(req)
			if err_res != nil {
				t.Fatal(err_res)
			}
			if res.StatusCode != 404 {
				t.Fatal("status code must be a 404", res.StatusCode)
			}

			// test with user_2's auth with organizer permissions
			req = httptest.NewRequest("GET", "/event/" + fmt.Sprint(user2_event.ID) + "/another-route", nil)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user_2_jwt))
			res, err_res = server.Test(req)
			if err_res != nil {
				t.Fatal(err_res)
			}
			if res.StatusCode != 200 {
				t.Fatal("status code must be a 200", res.StatusCode)
			}
			var body map[string]int64
			if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
				t.Fatal("could not parse body", err.Error())
			}
			if body["event_id"] != user2_event.ID {
				t.Fatal("event id is incorrect", body["event_id"], user2_event.ID)
			}
		})
	})
}
