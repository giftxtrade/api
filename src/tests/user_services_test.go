package tests

import (
	"fmt"
	"testing"

	"github.com/giftxtrade/api/src/app"
	"github.com/giftxtrade/api/src/services"
	"github.com/giftxtrade/api/src/types"
	"github.com/google/uuid"
)

func test_setup(t *testing.T) (*services.UserService) {
	db, err := NewMockDB(t)
	if err != nil {
		t.FailNow()
	}

	if app.AutoMigrate(db) != nil {
		fmt.Println("shit")
		t.FailNow()
	}

	db.Exec("delete from users")

	return &services.UserService{
		Service: services.Service{
			DB: db,
			TABLE: "users",
		},
	}
}

func TestCreateUser(t *testing.T) {
	user_service := test_setup(t)

	expected := types.CreateUser{
		Email: "john_doe@email.com",
		Name: "John Doe",
		ImageUrl: "https://images.com/john_doe",
	}

	t.Run("should create user", func(t *testing.T) {
		new_user, err := user_service.Create(&expected)

		if err != nil {
			fmt.Println("should not return an error", new_user, expected)
			t.Fail()
		}

		if new_user.ID == uuid.Nil || new_user.Name != expected.Name || new_user.Email != expected.Email || new_user.ImageUrl != expected.ImageUrl || !new_user.IsActive || new_user.IsAdmin {
			fmt.Println("user service create did not work", new_user, expected)
			t.Fail()
		}
	})

	t.Run("should not create new user", func(t *testing.T) {
		if _, err := user_service.Create(&expected); err == nil {
			fmt.Println("should not create a new user")
			t.Fail()
		}
	})
}