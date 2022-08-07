package tests

import (
	"testing"
	"time"

	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
)

func TestEventService(t *testing.T) {
	db := SetupMockProductService(t)
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
		event_input := types.CreateEvent{
			Name: "Event 1",
			Budget: 10,
			DrawAt: now,
			CloseAt: now,
		}
		created_event := types.Event{}
		err := event_service.Create(&event_input, &my_user, &created_event)
		if err != nil {
			t.Fatal("could not create event", err, event_input)
		}
	})
}