package user

import (
	"net/http"
	user_srv "user_srv/proto/front_user"
	"web/utils"

	"github.com/micro/go-grpc"
	"golang.org/x/net/context"

	"github.com/gin-gonic/gin"
)

func SendEmail(ctx *gin.Context) {
	email := ctx.PostForm("email")
	if utils.VerifyEmailFormat(email) {
		service := grpc.NewService()
		front_user_service := user_srv.NewFrontUserService("xlf.user.srv.user_srv", service.Client())
		rep, _ := front_user_service.FrontUserSendEmail(context.TODO(), &user_srv.FrontUserMailRequest{Email: email})

		ctx.JSON(http.StatusOK, gin.H{
			"code": rep.Code,
			"msg":  rep.Msg,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "邮箱格式不正确",
		})
	}
}

func FrontUserRegister(ctx *gin.Context) {
	email := ctx.PostForm("email")
	if utils.VerifyEmailFormat(email) {
		captche := ctx.PostForm("captche")
		password := ctx.PostForm("password")
		repassword := ctx.PostForm("repassword")

		if password != repassword {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "两次密码不一致",
			})
		} else {
			service := grpc.NewService()
			front_user_service := user_srv.NewFrontUserService("xlf.user.srv.user_srv", service.Client())
			rep, err := front_user_service.FrontUserRegister(
				context.TODO(), &user_srv.FrontUserRequest{
					Email:      email,
					Code:       captche,
					Password:   password,
					Repassword: repassword,
				})
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code": rep.Code,
					"msg":  rep.Msg,
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code": rep.Code,
					"msg":  rep.Msg,
				})
			}

		}

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "邮箱格式不正确",
		})
	}
}

func FrontUserLogin(ctx *gin.Context) {
	mail := ctx.PostForm("mail")
	password := ctx.PostForm("password")

	if utils.VerifyEmailFormat(mail) {
		service := grpc.NewService()
		front_user_service := user_srv.NewFrontUserService("xlf.user.srv.user_srv", service.Client())
		rep, err := front_user_service.FrontUserLogin(context.TODO(), &user_srv.FrontUserRequest{
			Email:    mail,
			Password: password,
		})

		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "用户名或密码错误",
			})
		} else {
			//生成token
			token, err2 := utils.GenToken(rep.UserName, utils.FrontUserExpireDuration, utils.FrontUserSecretKey)
			if err2 != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"code": rep.Code,
					"msg":  rep.Msg,
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"code":     rep.Code,
					"msg":      rep.Msg,
					"token":    token,
					"username": rep.UserName,
				})
			}

		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "邮箱地址不正确",
		})
	}

}
