package tests

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"maps"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/giftxtrade/api/src/controllers"
	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
	"github.com/giftxtrade/api/src/utils"
	"github.com/gofiber/fiber/v2"
)

func TestParticipant(t *testing.T) {
	token := app.Tokens.JwtKey
	user_service := app.Service.UserService
	event_service := app.Service.EventService

	// User #1
	user1, _, err := app.Service.UserService.FindOrCreate(context.Background(), types.CreateUser{
		Name: "Participant User #1",
		Email: "testparticipantuser1@email.com",
	})
	user1_jwt, _ := user_service.GenerateJWT(token, &user1)
	if err != nil {
		t.Fatal(err)
	}

	// User #2
	user2, _, err := app.Service.UserService.FindOrCreate(context.Background(), types.CreateUser{
		Name: "Participant User #2",
		Email: "testparticipantuser2@email.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	// user2_dto := mappers.DbUserToUser(user2)
	// user2_jwt, _ := user_service.GenerateJWT(token, &user2)

	// Event for user1
	event1_input := types.CreateEvent{
		Name: "Event #1 For User #1",
		Budget: 100.00,
		DrawAt: time.Now(),
		CloseAt: time.Now().Add(time.Hour * 24 * 30),
		Participants: append(create_participants(5), types.CreateParticipant{
			Name: user1.Name,
			Email: user1.Email,
			Organizer: true,
			Participates: false,
		}),
	}
	user1_event, err := event_service.CreateEvent(context.Background(), &user1, event1_input)
	if err != nil {
		t.Fatal(err)
	}
	// Event for user2
	// input := types.CreateEvent{
	// 	Name: "Event #2 For User #2",
	// 	Budget: 100.00,
	// 	DrawAt: time.Now(),
	// 	CloseAt: time.Now().Add(time.Hour * 24 * 30),
	// 	Participants: append(create_participants(5), types.CreateParticipant{
	// 		Name: user2.Name,
	// 		Email: user2.Email,
	// 		Organizer: true,
	// 		Participates: false,
	// 	}),
	// }
	// user2_event, err := event_service.CreateEvent(context.Background(), &user2_dto, input)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	t.Run("Participant middleware tests", func(t *testing.T) {
		t.Run("UseEventParticipantAuthWithQuery", func(t *testing.T) {
			var user1_participant types.Participant
			for _, p := range user1_event.Participants {
				if p.Email != user1.Email {
					continue
				}
				user1_participant = p
			}
			server.Get("/events/:event_id/manage", controller.UseJwtAuth, controller.UseEventAuthWithParam, controller.UseEventParticipantAuthWithQuery, func(c *fiber.Ctx) error {
				participant := c.UserContext().Value(controllers.PARTICIPANT_OB_KEY).(database.Participant)
				return utils.DataResponse(c, map[string]interface{}{"event_id": participant.EventID, "participant_id": participant.ID, "email": participant.Email})
			})

			// Unparsable participant id
			req := httptest.NewRequest("GET", fmt.Sprintf("/events/%d/manage?participantId=%s", user1_event.ID, "abc123"), nil)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user1_jwt))
			res, err_res := server.Test(req)
			if err_res != nil {
				t.Fatal(err_res)
			}
			if res.StatusCode != 400 {
				t.Fatal("status code must be a 400", res.StatusCode)
			}

			// Unknown participant id
			req = httptest.NewRequest("GET", fmt.Sprintf("/events/%d/manage?participantId=%d", user1_event.ID, 3483), nil)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user1_jwt))
			res, err_res = server.Test(req)
			if err_res != nil {
				t.Fatal(err_res)
			}
			if res.StatusCode != 400 {
				t.Fatal("status code must be a 400", res.StatusCode)
			}
			var body map[string]([]string)
			if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
				t.Fatal(err)
			}
			if body["errors"][0] != "participant does not exist on the event" {
				t.Fatal(body)
			}

			// Correct participant ID
			req = httptest.NewRequest("GET", fmt.Sprintf("/events/%d/manage?participantId=%d", user1_event.ID, user1_participant.ID), nil)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user1_jwt))
			res, err_res = server.Test(req)
			if err_res != nil {
				t.Fatal(err_res)
			}
			if res.StatusCode != 200 {
				t.Fatal("status code must be a 200", res.StatusCode)
			}
			var body2 map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&body2); err != nil {
				t.Fatal(err)
			}
			data := map[string]interface{}{"event_id": user1_event.ID, "participant_id": user1_participant.ID, "email": user1_participant.Email}
			if maps.Equal(body2, data) {
				t.Fatal(body2, data)
			}

			// Correct participant id with
			user2_participant, err := app.Querier.CreateParticipant(context.Background(), database.CreateParticipantParams{
				EventID: user1_event.ID,
				UserID: sql.NullInt64{
					Valid: true,
					Int64: user2.ID,
				},
				Accepted: true,
				Participates: true,
				Name: user2.Name,
				Email: user2.Email,
			})
			if err != nil {
				t.Fatal(err)
			}
			req = httptest.NewRequest("GET", fmt.Sprintf("/events/%d/manage?participantId=%d", user1_event.ID, user2_participant.ID), nil)
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", user1_jwt))
			res, err_res = server.Test(req)
			if err_res != nil {
				t.Fatal(err_res)
			}
			if res.StatusCode != 200 {
				t.Fatal("status code must be a 200", res.StatusCode)
			}
			if err := json.NewDecoder(res.Body).Decode(&body2); err != nil {
				t.Fatal(err)
			}
			data = map[string]interface{}{"event_id": user1_event.ID, "participant_id": user2_participant.ID, "email": user2_participant.Email}
			if maps.Equal(body2, data) {
				t.Fatal(body2, data)
			}
		})
	})
}
