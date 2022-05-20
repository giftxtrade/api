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
		t.Fatal("migration failed")
	}

	db.Exec("delete from users")

	return &services.UserService{
		Service: services.Service{
			DB: db,
			TABLE: "users",
		},
	}
}

func TestUserService(t *testing.T) {
	user_service := test_setup(t)

	test_user1 := types.CreateUser{
		Email: "john_doe@email.com",
		Name: "John Doe",
		ImageUrl: "https://images.com/john_doe",
	}

	test_user2 := &types.CreateUser{
		Name: "Test User",
		Email: "testuser@email.com",
		ImageUrl: "https://images.com/test_user2",
	}

	t.Run("should create user", func(t *testing.T) {
		new_user, err := user_service.Create(&test_user1)

		if err != nil {
			fmt.Println("should not return an error", new_user, test_user1)
			t.Fail()
		}

		if new_user.ID == uuid.Nil || new_user.Name != test_user1.Name || new_user.Email != test_user1.Email || new_user.ImageUrl != test_user1.ImageUrl || !new_user.IsActive || new_user.IsAdmin {
			fmt.Println("user service create did not work", new_user, test_user1)
			t.Fail()
		}
	})

	t.Run("should not create new user", func(t *testing.T) {
		if _, err := user_service.Create(&test_user1); err == nil {
			fmt.Println("should not create a new user")
			t.Fail()
		}
	})

	t.Run("should find user", func(t *testing.T) {
		t.Run("should find by email", func(t *testing.T) {
			user_by_email, err := user_service.FindByEmail(test_user1.Email)
			if err != nil {
				t.Fatal(err)
			}
			if user_by_email.Email != test_user1.Email || user_by_email.Name != test_user1.Name || user_by_email.ID == uuid.Nil {
				t.FailNow()
			}

			if _, err = user_service.FindByEmail(user_by_email.ID.String()); err == nil {
				t.Fatal(err)
			}
		})

		t.Run("should find by id", func(t *testing.T) {
			user_by_email, err := user_service.FindByEmail(test_user1.Email)
			if err != nil {
				t.Fatal(err)
			}
			user_by_id, err := user_service.FindById(user_by_email.ID.String())
			if err != nil {
				t.Fatal(err)
			}
			if user_by_id.Email != test_user1.Email || user_by_id.Name != test_user1.Name || user_by_id.ID == uuid.Nil {
				t.FailNow()
			}

			if _, err = user_service.FindById(user_by_id.Email); err == nil {
				t.Fatal(err)
			}
		})

		t.Run("should find or create", func(t *testing.T) {
			created_user, created, err := user_service.FindOrCreate(test_user2)
			if err != nil || !created {
				t.Fatal(err)
			}
			if created_user.Email != test_user2.Email || created_user.Name != test_user2.Name || created_user.ID == uuid.Nil {
				t.FailNow()
			}

			found_user, created, err := user_service.FindOrCreate(test_user2)
			if err != nil || created {
				t.Fatal(err)
			}
			if found_user.Email != test_user2.Email || found_user.Name != test_user2.Name || found_user.ID == uuid.Nil {
				t.FailNow()
			}
		})

		t.Run("should find with id and email", func(t *testing.T) {
			_, err := user_service.FindByIdAndEmail(uuid.NewString(), test_user1.Email)
			if err == nil {
				t.Fatal(err)
			}

			user_by_email, err := user_service.FindByEmail(test_user1.Email)
			if err != nil {
				t.Fatal(err)
			}

			user, err := user_service.FindByIdAndEmail(user_by_email.ID.String(), test_user1.Email)
			if err != nil || user.ID != user_by_email.ID || user.Email != user_by_email.Email {
				t.Fatal(err, user, user_by_email)
			}
		})
	})

	t.Cleanup(func (){
		user_service.DB.Exec(fmt.Sprintf("DELETE FROM %s", user_service.TABLE))
	})
}