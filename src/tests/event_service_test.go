package tests

import (
	"testing"
	"time"

	"github.com/giftxtrade/api/src/types"
)

func TestEventService(t *testing.T) {
	app := New(t)
	event_service := app.Service.EventService
	my_user := types.User{}
	_, user_create_err := event_service.UserService.FindOrCreate(
		&types.CreateUser{
			Name: "Event test user",
			Email: "event_test_user@giftxtrade.com",
		},
		&my_user,
	)
	if user_create_err != nil {
		t.Fatal(user_create_err)
	}

	t.Run("create event", func(t *testing.T) {
		now := time.Now()
		input := types.CreateEvent{
			Name: "Event 1",
			Budget: 10,
			DrawAt: now,
			CloseAt: now,
		}
		event := types.Event{}
		err := event_service.Create(&input, &my_user, &event)
		if err != nil {
			t.Fatal("could not create event", err, input)
		}
		if event.Name != input.Name || event.Budget != input.Budget || event.DrawAt != input.DrawAt || event.CloseAt != input.CloseAt || event.ModifiedById != event.CreatedById {
			t.Fatal("created event does not have values from input", event, input)
		}
		if event.CreatedBy.ID != my_user.ID || event.ModifiedBy.ID != my_user.ID {
			t.Fatal("incorrect event owner")
		}
	})

	t.Run("find event by id", func(t *testing.T) {
		now := time.Now()
		input := types.CreateEvent{
			Name: "Event 2",
			Budget: 6.99,
			DrawAt: now,
			CloseAt: now,
		}
		event := types.Event{}
		err := event_service.Create(&input, &my_user, &event)
		if err != nil {
			t.Fatal("could not create event", err, input)
		}

		event_by_id := types.Event{}
		found_err := event_service.FindById(event.ID.String(), &event_by_id)
		if found_err != nil {
			t.Fatal(found_err)
		}
		if event_by_id.ID != event.ID || event_by_id.Name != event.Name {
			t.Fatal("events not equal", event, event_by_id)
		}
	})

	t.Run("patch event", func(t *testing.T) {
		now := time.Now()
		input := types.CreateEvent{
			Name: "Event 2",
			Budget: 6.99,
			DrawAt: now,
			CloseAt: now,
		}
		var event types.Event
		if err := event_service.Create(&input, &my_user, &event); err != nil {
			t.Fatal(err)
		}

		t.Run("patch nothing", func(t *testing.T) {
			t.Run("default values", func(t *testing.T) {
				updated_event := event
				updated, err := event_service.Patch(&my_user, &types.CreateEvent{}, &event)
				if err != nil {
					t.Fatal(err)
				}
				if updated == true {
					t.Fatal("event should not update. all default values")
				}
				if updated_event.ModifiedBy.ID != event.ModifiedBy.ID {
					t.Fatal("modified by user should not be changed")
				}
			})

			t.Run("original event values", func(t *testing.T) {
				updated_event := event
				updated, err := event_service.Patch(&my_user, &input, &updated_event)
				if err != nil {
					t.Fatal(err)
				}
				if updated == true {
					t.Fatal("event should not update. values did not change", event)
				}
				if updated_event.ModifiedBy.ID != event.ModifiedBy.ID {
					t.Fatal("modified by user should not be changed")
				}
			})
		})

		t.Run("update values", func(t *testing.T) {
			input := types.CreateEvent{
				Name: "Event 2 (Updated)",
			}
			updated_event := event

			// create new user
			new_user_input := types.CreateUser{
				Email: "json@batman.com",
				Name: "Json Todd",
			}
			new_user := types.User{}
			_, user_create_err := event_service.UserService.FindOrCreate(&new_user_input, &new_user)
			if user_create_err != nil {
				t.Fatal("could not create new user", user_create_err)
			}

			updated, err := event_service.Patch(&new_user, &input, &updated_event)
			if err != nil {
				t.Fatal("could not patch event", err)
			}
			if !updated {
				t.Fatal("event was not updated", updated_event)
			}
			if updated_event.Name != input.Name {
				t.Fatal("event name was not updated", updated_event, input)
			}
			if updated_event.ID != event.ID {
				t.Fatal("event id should never update")
			}
			if updated_event.ModifiedById != new_user.ID {
				t.Fatal("modified user not assigned properly")
			}
		})
	})

	t.Run("delete event", func(t *testing.T) {
		now := time.Now()
		input := types.CreateEvent{
			Name: "Event to be deleted",
			Budget: 499.99,
			DrawAt: now,
			CloseAt: now,
			Description: "Some random even description",
		}
		var event types.Event
		if err := event_service.Create(&input, &my_user, &event); err != nil {
			t.Fatal(err)
		}
		event_id := event.ID.String()

		t.Run("valid event id", func(t *testing.T) {
			if err := event_service.Delete(event_id); err != nil {
				t.Fatal("should delete event with id", event_id)
			}
			found_event := types.Event{}
			if err := event_service.FindById(event_id, &found_event); err == nil {
				t.Fatal("event should have been deleted already")
			}
		})
	})

	t.Cleanup(func() {
		event_service.DB.Exec("delete from users, events")
	})
}