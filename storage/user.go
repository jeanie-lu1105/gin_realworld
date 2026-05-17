package storage

import (
	"context"

	"example.com/gin_realworld/logger"
	"example.com/gin_realworld/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	log := logger.New(ctx)
	_, err := db.ExecContext(ctx, "inset user(username, password, emial, image, bio) values (?,?,?,?,?)", user.Username, user.Password, user.Email, user.Image, user.Bio)
	log.Infof("create user in db but error ", err)
	return err
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := db.GetContext(ctx, &user, "select * from user where email = ?", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := db.GetContext(ctx, &user, "select * from user where username = ?", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := db.ExecContext(ctx, "delete from user where email = ?", email)
	return err
}

func UpdateUserByUsername(ctx context.Context, username string, user *models.User) error {
	if user.Password != "" {
		_, err := db.ExecContext(ctx, "update user set username=?, password=?, email=?, image=?,bio=? where username=?", user.Username, user.Password, user.Email, user.Image, user.Bio, username)
		return err
	}
	_, err := db.ExecContext(ctx, "update user set username=?, email=?, image=?, bio=? where username =?", user.Username, user.Email, user.Image, user.Bio, username)
	return err
}
