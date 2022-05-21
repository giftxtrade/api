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

	t.Run("create user", func(t *testing.T) {
		t.Run("should create user", func(t *testing.T) {
			new_user, err := user_service.Create(&test_user1)

			if err != nil {
				t.Fatal("should not return an error", new_user, test_user1)
			}

			if new_user.ID == uuid.Nil || new_user.Name != test_user1.Name || new_user.Email != test_user1.Email || new_user.ImageUrl != test_user1.ImageUrl || !new_user.IsActive || new_user.IsAdmin {
				t.Fatal("user service create did not work", new_user, test_user1)
			}
		})

		t.Run("should not create existing user", func(t *testing.T) {
			if _, err := user_service.Create(&test_user1); err == nil {
				t.Fatalf("should not create a new user")
			}
		})
	})

	t.Run("find user", func(t *testing.T) {
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
				t.Fatal("should not find a user with an empty uuid")
			}

			user_by_email, err := user_service.FindByEmail(test_user1.Email)
			if err != nil {
				t.Fatal(err)
			}

			user, err := user_service.FindByIdAndEmail(user_by_email.ID.String(), test_user1.Email)
			if err != nil || user.ID != user_by_email.ID || user.Email != user_by_email.Email {
				t.Fatal(err, user, user_by_email)
			}

			_, err = user_service.FindByIdAndEmail(user.ID.String(), test_user2.Email)
			if err == nil {
				t.Fatal("should not find a user with id and email from different users")
			}
		})

		t.Run("should find user with any primary key", func(t *testing.T) {
			if _, err := user_service.Find("random key"); err == nil {
				t.Fatal("user should not exist")
			}
			if _, err := user_service.Find(uuid.NewString()); err == nil {
				t.Fatal("user should not exist")
			}

			test_user, _ := user_service.FindByEmail(test_user2.Email)

			u1, err1 := user_service.Find(test_user.Email)
			u2, err2 := user_service.Find(test_user.ID.String())
			if err1 != nil || err2 != nil {
				t.Fatal(err1, err2)
			}
			if u1.Email != u2.Email || u1.ID != u2.ID {
				t.Fatal("both users should match")
			}
		})
	})

	t.Cleanup(func() {
		user_service.DB.Exec(fmt.Sprintf("DELETE FROM %s", user_service.TABLE))
	})
}