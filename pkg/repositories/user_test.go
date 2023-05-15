package repositories

import (
	"go-blog/pkg/models"
	"log"
	"testing"
)

var registrationDemoUsers = []models.User{
	{
		Username: "johndoe",
		Email:    "john@example.com",
		Password: "12345678"},
	{
		Username: "janebou",
		Email:    "jane@example.com",
		Password: "12345678"},
	{
		Username: "bobchele",
		Email:    "bob@example.com",
		Password: "12345678"},
	{
		Username: "user1",
		Email:    "user1@example.com",
		Password: "12345678",
	},
	{
		Username: "user2",
		Email:    "user2@example.com",
		Password: "12345678",
	},
}

func TestRegistration(t *testing.T) {

	for _, newUser := range registrationDemoUsers {
		err := db.Create(&newUser).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

}
