package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/giftxtrade/api/src/database"
	"github.com/giftxtrade/api/src/types"
)

func TestUserService(t *testing.T) {
	app := New(t)
	user_service := app.Service.UserService

	test_user1 := database.CreateUserParams{
		Email: "john_doe@email.com",
		Name: "John Doe",
		ImageUrl: "https://images.com/john_doe",
	}

	test_user2 := database.CreateUserParams{
		Name: "Test User",
		Email: "testuser@email.com",
		ImageUrl: "https://images.com/test_user2",
	}

	t.Run("create user", func(t *testing.T) {
		t.Run("should create user", func(t *testing.T) {
			new_user, err := user_service.Querier.CreateUser(context.Background(), test_user1)
			if err != nil {
				t.Fatal("should not return an error", new_user, test_user1)
			}

			if new_user.ID == 0 || new_user.Name != test_user1.Name || new_user.Email != test_user1.Email || new_user.ImageUrl != test_user1.ImageUrl || new_user.Active || new_user.Admin {
				t.Fatal("user service create did not work", new_user, test_user1)
			}
		})

		t.Run("should not create existing user", func(t *testing.T) {
			if _, err := user_service.Querier.CreateUser(context.Background(), test_user1); err == nil {
				t.Fatalf("should not create a new user")
			}
		})
	})

	t.Run("find user", func(t *testing.T) {
		t.Run("should find by email", func(t *testing.T) {
			user_by_email, err := user_service.Querier.FindUserByEmail(context.Background(), test_user1.Email)
			if err != nil {
				t.Fatal(err)
			}
			if user_by_email.Email != test_user1.Email || user_by_email.Name != test_user1.Name || user_by_email.ID == 0 {
				t.FailNow()
			}

			if user_by_email, err = user_service.Querier.FindUserByEmail(context.Background(), fmt.Sprint(user_by_email.ID)); err == nil {
				t.Fatal(err)
			}
		})

		t.Run("should find by id", func(t *testing.T) {
			user_by_email, err := user_service.Querier.FindUserByEmail(context.Background(), test_user1.Email)
			if err != nil {
				t.Fatal(err)
			}

			user_by_id, err := user_service.Querier.FindUserById(context.Background(), user_by_email.ID)
			if err != nil {
				t.Fatal(err)
			}
			if user_by_id.Email != test_user1.Email || user_by_id.Name != test_user1.Name || user_by_id.ID == 0 {
				t.FailNow()
			}
		})

		t.Run("should find or create", func(t *testing.T) {
			created_user, created, err := user_service.FindOrCreate(context.Background(), types.CreateUser{
				Name: test_user2.Name,
				Email: test_user2.Email,
				ImageUrl: test_user2.ImageUrl,
			})
			if err != nil || !created {
				t.Fatal(err)
			}
			if created_user.Email != test_user2.Email || created_user.Name != test_user2.Name || created_user.ID == 0 {
				t.FailNow()
			}

			found_user, created, err := user_service.FindOrCreate(context.Background(), types.CreateUser{
				Name: test_user2.Name,
				Email: test_user2.Email,
				ImageUrl: test_user2.ImageUrl,
			})
			if err != nil || created {
				t.Fatal(err)
			}
			if found_user.Email != test_user2.Email || found_user.Name != test_user2.Name || found_user.ID != created_user.ID {
				t.FailNow()
			}
		})

		t.Run("should find with id and email", func(t *testing.T) {	
			user, err := user_service.Querier.FindUserByIdAndEmail(context.Background(), database.FindUserByIdAndEmailParams{
				ID: 3434,
				Email: test_user1.Email,
			})
			if err == nil {
				t.Fatal("should not find a user with an non existing or matching uuid")
			}

			user_by_email, err := user_service.Querier.FindUserByEmail(context.Background(), test_user1.Email)
			if err != nil {
				t.Fatal(err)
			}
 
			user, err = user_service.Querier.FindUserByIdAndEmail(context.Background(), database.FindUserByIdAndEmailParams{
				ID: user_by_email.ID,
				Email: test_user1.Email,
			})
			if err != nil || user.ID != user_by_email.ID || user.Email != user_by_email.Email {
				t.Fatal(err, user, user_by_email)
			}
			
			_, err = user_service.Querier.FindUserByIdAndEmail(context.Background(), database.FindUserByIdAndEmailParams{
				ID: user.ID,
				Email: test_user2.Email,
			})
			if err == nil {
				t.Fatal("should not find a user with id and email from different users")
			}
		})
	})
}