package tests

import (
	"testing"
	"time"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
)

func TestEventService(t *testing.T) {
	db := MockMigration(t)
	event_service := services.EventService{
		ServiceBase: services.CreateService(db, "events"),
		UserService: services.UserService{
			ServiceBase: services.CreateService(db, "users"),
		},
	}
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
			original_event := event
			t.Run("default values", func(t *testing.T) {
				updated, err := event_service.Patch(event.ID.String(), &my_user, &types.CreateEvent{}, &event)
				if err != nil {
					t.Fatal(err)
				}
				if updated == true {
					t.Fatal("event should not update. all default values")
				}
				if event.ModifiedBy.ID != original_event.ModifiedBy.ID {
					t.Fatal("modified by user should not be changed")
				}
			})

			t.Run("original event values", func(t *testing.T) {
				updated, err := event_service.Patch(event.ID.String(), &my_user, &input, &event)
				if err != nil {
					t.Fatal(err)
				}
				if updated == true {
					t.Fatal("event should not update. values did not change", event)
				}
				if event.ModifiedBy.ID != original_event.ModifiedBy.ID {
					t.Fatal("modified by user should not be changed")
				}
			})
		})
	})

	t.Cleanup(func() {
		event_service.DB.Exec("delete from users, events")
	})
}