package storage

import (
	"context"
	"testing"

	"example.com/gin_realworld/models"
	"example.com/gin_realworld/utils"
)

func TestCreateAndGetUser(t *testing.T) {
	ctx := context.Background()
	userName := "jessie_lu"
	email := "jessie_lu@gmail.com"

	err := CreateUser(ctx, &models.User{
		Username: userName,
		Password: "12345678",
		Email:    email,
		Image:    "https://image.baidu.com/front/aigc?atn=aigc&fr=home_hover&imgcontent=%7B%22aigcQuery%22%3A%22%22%2C%22imageAigcId%22%3A%22%22%7D&isImmersive=1&pd=image_content&quality=1&ratio=3%3A4&sa=searchpromo_shijian_photohp_inspire&tn=aigc&top=%7B%22sfhs%22%3A1%7D&word=%E5%90%91%E6%97%A5%E8%91%B5%E8%8A%B1%E4%B8%9B%E4%B8%AD%EF%BC%8C%E6%8A%B1%E7%9D%80%E5%90%91%E6%97%A5%E8%91%B5%E7%9A%84%E5%B9%B4%E8%BD%BB%E5%A5%B3%E5%AD%A9%EF%BC%8C%E5%BC%80%E5%BF%83%E5%9C%B0%E5%BE%AE%E7%AC%91%EF%BC%8C%E6%9A%96%E9%87%91%E8%89%B2%E7%9A%84%E9%98%B3%E5%85%89",
		Bio:      "xxxx123",
	})
	if err != nil {
		t.Errorf("create user failed, err: %v", err)
		return
	}

	dbUser, err := GetUserByEmail(ctx, email)
	if err != nil {
		t.Errorf("get user by email failed, err: %v", err)
		return
	}

	t.Logf("get user: %v\n", utils.JsonMarshal(dbUser))

	err = DeleteUserByEmail(ctx, email)
	if err != nil {
		t.Errorf("delete user by email failed")
		return
	}

}
