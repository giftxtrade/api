package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/giftxtrade/api/src/types"
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
	app := New(t)
	event_service := app.Service.EventService
	user, _, err := app.Service.UserService.FindOrCreate(context.Background(), types.CreateUser{
		Name: "Test User",
		Email: "testuser@email.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	user_dto := types.User{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Active: user.Active,
	}

	t.Run("create event", func(t *testing.T) {
		t.Run("correct params", func(t *testing.T) {
			event := types.CreateEvent{
				Name: "Event 1",
				Budget: 100.00,
				DrawAt: time.Now(),
				CloseAt: time.Now().Add(time.Hour * 24 * 30),
				Participants: append(create_participants(5), types.CreateParticipant{
					Name: user.Name,
					Email: user.Email,
					Organizer: true,
					Participates: false,
				}),
			}
			new_event, err := event_service.CreateEvent(context.Background(), &user_dto, event)
			if err != nil {
				t.Fatal(err)
			}

			if len(new_event.Participants) != 6 {
				t.Fatal("not all participants were inserted")
			}
		})

		t.Run("main participant not marked as organizer", func(t *testing.T) {
			event := types.CreateEvent{
				Name: "Event 1",
				Budget: 100.00,
				DrawAt: time.Now(),
				CloseAt: time.Now().Add(time.Hour * 24 * 30),
				Participants: append(create_participants(5), types.CreateParticipant{
					Name: user.Name,
					Email: user.Email,
					Organizer: false,
					Participates: true,
				}),
			}
			_, err := event_service.CreateEvent(context.Background(), &user_dto, event)
			if err == nil {
				t.Fatal("event should not be created. main participant is not marked 'organizer'")
			}

			event.Participants = create_participants(20)
			_, err = event_service.CreateEvent(context.Background(), &user_dto, event)
			if err == nil {
				t.Fatal("event should not be created. main participant was not provided")
			}
		})
	})
}
