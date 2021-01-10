package controller

import (
	"context"
	"fmt"
	"time"
	"user_srv/data_source"
	"user_srv/models"
	user_srv "user_srv/proto/front_user"
	"user_srv/utils"

	"github.com/pkg/errors"

	"github.com/patrickmn/go-cache"
)

type FrontUser struct{}

var c = cache.New(300*time.Second, 10*time.Second)

func (f *FrontUser) FrontUserRegister(ctx context.Context, in *user_srv.FrontUserRequest, out *user_srv.FrontUserRespond) error {
	//用户注册

	email := in.Email
	captche := in.Code
	password := in.Password

	//验证码是否正确
	code, is_ok := c.Get(email)
	if is_ok {
		if code != captche {
			out.Code = 500
			out.Msg = "邮箱验证码错误"
		} else {
			//保存数据到数据库
			md5_password := utils.Md5pwd(password)
			front_user := models.FrontUser{
				Email:      email,
				Password:   md5_password,
				Status:     1,
				CreateTime: time.Now(),
			}
			data_source.Db.Create(&front_user)
			out.Code = 200
			out.Msg = "注册成功，请登录"
		}
	} else {
		out.Code = 500
		out.Msg = "邮箱验证码错误"
	}
	return nil
}

func (f *FrontUser) FrontUserSendEmail(ctx context.Context, in *user_srv.FrontUserMailRequest, out *user_srv.FrontUserRespond) error {
	email := in.Email
	//六位验证码

	front_user := models.FrontUser{}
	var count int
	data_source.Db.Where("email=?", email).Find(&front_user).Count(&count)
	fmt.Println("------------------------")
	fmt.Println(count)
	if count == 0 {
		rand_num := utils.GenRandNum(6)
		utils.SendEmail(email, rand_num)
		//session
		c.Set(email, rand_num, cache.DefaultExpiration)
		out.Code = 200
		out.Msg = "发送成功"
	} else {
		out.Code = 500
		out.Msg = "邮箱已被注册"
	}
	return nil
}

func (f *FrontUser) FrontUserLogin(ctx context.Context, in *user_srv.FrontUserRequest, out *user_srv.FrontUserRespond) error {
	email := in.Email
	password := in.Password
	md5_password := utils.Md5pwd(password)

	front_user := models.FrontUser{}
	var count int
	data_source.Db.Where("email=?", email).Where("password=?", md5_password).Find(&front_user).Count(&count)
	if count < 1 {
		out.Code = 500
		out.Msg = "用户名或密码错误"
		return errors.New("用户名或密码错误")
	} else {
		out.Code = 200
		out.Msg = "登录成功"
		out.UserName = email
		return nil
	}
}
