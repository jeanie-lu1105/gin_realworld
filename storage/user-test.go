package storage

import (
	"context"
	"testing"

	"example.com/gin_realworld/models"
	"example.com/gin_realworld/utils"
)

func TestCreateAndGetUser(t *testing.T) {
	ctx := context.Background()
	userName := "lujiaxinTest"
	email := "lujiaxin1@test.com"

	err := CreateUser(ctx, &models.User{
		Username: userName,
		Password: "xxxx123",
		Email:    email,
		Image:    "xxxx123",
		Bio:      "xxxx123",
	})

	if err != nil {
		t.Errorf("CreateUser err: %v", err)
		return
	}

	err = CreateUser(ctx, &models.User{
		Username: "lujiaxin_test",
		Password: "xxxx1234",
		Email:    "lujiaxin_test@test.com",
		Image:    "xxxx123",
		Bio:      "xxxx123",
	})

	if err != nil {
		t.Errorf("CreateUser err: %v", err)
		return
	}

	dbUser, err := GetUserByEmail(ctx, email)
	if err != nil {
		t.Errorf("GetUserByEmail err: %v", err)
		return
	}

	t.Logf("dbUser: %v", utils.JsonMarshal(dbUser))

	err = DeleteUserByEmail(ctx, email)
	if err != nil {
		t.Errorf("DeleteUserByEmail err: %v", err)
		return
	}
}
