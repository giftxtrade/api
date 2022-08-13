package tests

import (
	"testing"
	"time"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
)

func TestEventService(t *testing.T) {
	db := SetupMockEventService(t)
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
		if err != nil {
			t.Fatal("could not create event", err, event_input)
		}
	})

	t.Cleanup(func() {
		event_service.DB.Exec("delete from users, events")
	})
}