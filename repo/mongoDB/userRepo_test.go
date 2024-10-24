package mongodb

import (
	"testing"

	"github.com/google/uuid"
	"github.com/nilspolek/Workout-Tracker/repo"
)

func TestNewUserRepository(t *testing.T) {
	userRepo, err := NewUserRepository("mongodb://admin:admin123@localhost:27017")
	if err != nil {
		t.Errorf("Error creating new user repository: %v", err)
	}
	if userRepo == nil {
		t.Errorf("User repository is nil")
	}
	err = userRepo.CreateUser(repo.User{
		Id:       uuid.New(),
		Username: "test",
		Password: "test",
		Email:    "abc@def.gh",
	})
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}
	user, err := userRepo.GetUser("test")
	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}
	if user.Username != "test" {
		t.Errorf("Username is not test its: %s", user.Username)
	}
	err = userRepo.Close()
	if err != nil {
		t.Errorf("Error closing user repository: %v", err)
	}
}

func TestJWT(t *testing.T) {
	userRepo, err := NewUserRepository("mongodb://admin:admin123@localhost:27017")
	if err != nil {
		t.Errorf("Error creating new user repository: %v", err)
	}
	token, err := userRepo.GenerateJWT(uuid.New())
	if err != nil {
		t.Errorf("Error generating jwt: %v", err)
	}
	valid, err := userRepo.ValidateJWT(token)
	if err != nil {
		t.Errorf("Error validating jwt: %v", err)
	}
	if !valid {
		t.Errorf("Token is not valid")
	}
	if err != nil {
		t.Errorf("Error closing user repository: %v", err)
	}
	valid, err = userRepo.ValidateJWT("asd;klasd;lk")
	if err != nil {
		t.Errorf("Error validating jwt: %v", err)
	}
	if valid {
		t.Errorf("Token is valid but should not be")
	}
	err = userRepo.Close()
}
