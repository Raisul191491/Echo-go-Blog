package services

import (
	"fmt"
	"go-blog/pkg/connection"
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

	fmt.Println("fsdfsfsdd")
	testDB := connection.GetTestDB()

	for _, newUser := range registrationDemoUsers {
		err := testDB.Create(&newUser).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

}
