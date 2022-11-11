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

		t.Run("invalid event id", func(t *testing.T) {
			input := types.CreateParticipant{
				Email: "my_test_email@giftxtrade.com",
				Organizer: false,
				Participates: true,
			}
			participant := types.Participant{}

			err := participant_service.Create(
				&my_user, 
				nil, 
				&types.Event{
					Base: types.Base {
						ID: uuid.New(),
					},
				}, 
				&input, 
				&participant,
			)
			if err == nil {
				t.Fatal("event id is invalid. should return an error")
			}
		})

		t.Run("participant_user", func(t *testing.T) {
			const EMAIL = "my_test_email_2@giftxtrade.com"
			test_user := types.User{}
			_, user_create_err := app.Service.UserService.FindOrCreate(
				&types.CreateUser{
					Name: "Valid Participant User",
					Email: EMAIL,
				},
				&test_user,
			)
			if user_create_err != nil {
				t.Fatal("could not create user")
			}

			t.Run("valid emails", func(t *testing.T) {
				input := types.CreateParticipant{
					Email: EMAIL,
					Organizer: true,
					Participates: true,
					Nickname: "TopG",
					Address: "123 South Randall St.",
				}
				participant := types.Participant{}
	
				err := participant_service.Create(&my_user, &test_user, &event, &input, &participant)
				if err != nil {
					t.Fatal("could not create participant", err)
				}
	
				if participant.Event.ID != event.ID {
					t.Fatal("incorrect event id", participant.Event, event)
				}
				if !participant.UserId.Valid {
					t.Fatal("user must be inserted")
				}
				if participant.UserId.UUID != test_user.ID {
					t.Fatal("user ids don't match", participant.UserId, test_user.ID)
				}
				if !participant.Accepted {
					t.Fatal("participant must be accepted")
				}

				check := types.CreateParticipant{
					Email: participant.Email,
					Address: participant.Address,
					Nickname: participant.Nickname,
					Organizer: participant.Organizer,
					Participates: participant.Participates,
				}
				if check != input {
					t.Fatal("wrong values")
				}
			})
	
			t.Run("invalid emails", func(t *testing.T) {
				input := types.CreateParticipant{
					Email: "some_random_email@giftxtrade.com",
					Organizer: false,
					Participates: true,
				}
				participant := types.Participant{}
	
				err := participant_service.Create(&my_user, &test_user, &event, &input, &participant)
				if err == nil {
					t.Fatal("emails don't match")
				}
			})
		})

	})

	t.Run("find participant", func(t *testing.T) {
		input := types.CreateParticipant{
			Email: "find_particpant_test@giftxtrade.com",
			Organizer: true,
			Participates: true,
		}
		participant := types.Participant{}
		create_err := participant_service.Create(&my_user, nil, &event, &input, &participant)
		if create_err != nil {
			t.Fatal("could not create participant", create_err)
		}

		findTest := func(t *testing.T, result *types.Participant) {
			if result.ID != participant.ID {
				t.Fatal("incorrect event id")
			}
			if result.Email != participant.Email {
				t.Fatal("incorrect email")
			}

			// test for correct joins
			if result.ModifiedById != participant.ModifiedById && result.ModifiedBy != participant.ModifiedBy {
				t.Fatal("incorrect join for ModifiedBy field", result, participant)
			}
			if result.CreatedById != participant.CreatedById && result.CreatedBy != participant.CreatedBy {
				t.Fatal("incorrect join for CreatedBy field", result, participant)
			}
			if result.Event.ID != event.ID {
				t.Fatal("incorrect join for event", result, participant)
			}
		}

		t.Run("find by id", func(t *testing.T) {
			result := types.Participant{}
			err := participant_service.FindById(participant.ID.String(), &result)
			if err != nil {
				t.Fatal("could not find participant", err)
			}

			findTest(t, &result)
		})

		t.Run("find by email and event id", func(t *testing.T) {
			result := types.Participant{}
			err := participant_service.Find(participant.Email, event.ID.String(), &result)
			if err != nil {
				t.Fatal("could not find participant", err)
			}

			findTest(t, &result)
		})
	})

	t.Cleanup(func() {
		event_service.DB.Exec("delete from participants, events, users")
	})
}