package tests

import (
	"testing"

	"github.com/giftxtrade/api/src/types"
	"github.com/google/uuid"
)

func TestParticipantService(t *testing.T) {
	app := New(t)
	participant_service := app.Service.ParticipantService
	event_service := app.Service.EventService

	my_user := types.User{}
	_, user_create_err := event_service.UserService.FindOrCreate(
		&types.CreateUser{
			Name: "Participant test user",
			Email: "participant_test_user@giftxtrade.com",
		},
		&my_user,
	)
	if user_create_err != nil {
		t.Fatal(user_create_err)
	}

	event := types.Event{}
	event_create_err := event_service.Create(
		&types.CreateEvent{
			Name: "My new event",
			Description: "Participant test event",
			Budget: 39.99,
			DrawAt: GetTomorrow(),
			CloseAt: GetTomorrow(),
		},
		&my_user,
		&event,
	)
	if event_create_err != nil {
		t.Fatal("could not create event", event_create_err)
	}

	t.Run("create participant", func(t *testing.T) {
		t.Run("valid input", func(t *testing.T) {
			input := types.CreateParticipant{
				Email: "my_test_email@giftxtrade.com",
				Organizer: false,
				Participates: true,
			}
			participant := types.Participant{}

			err := participant_service.Create(&my_user, nil, &event, &input, &participant)
			if err != nil {
				t.Fatal("could not create participant", err)
			}

			if participant.Event.ID != event.ID {
				t.Fatal("incorrect event id", participant.Event, event)
			}
			if participant.UserId.Valid && participant.UserId.UUID != uuid.Nil {
				t.Fatal("user must not be defined")
			}
		})
	})

	t.Cleanup(func() {
		event_service.DB.Exec("delete from participants, events, users")
	})
}