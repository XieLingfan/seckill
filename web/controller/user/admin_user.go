package user

import (
	"net/http"
	user_srv "user_srv/proto/admin_user"
	"web/utils"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-grpc"
	"golang.org/x/net/context"
)

func AdminLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	service := grpc.NewService()
	admin_user_service := user_srv.NewAdminUserService("xlf.user.srv.user_srv", service.Client())
	rep, err := admin_user_service.AdminUserLogin(context.TODO(), &user_srv.AdminUserRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "用户名或密码错误",
		})
	} else {
		//生成token
		token, err2 := utils.GenToken(rep.Username, utils.AdminUserExpireDuration, utils.AdminUserSecretKey)
		if err2 != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": rep.Code,
				"msg":  rep.Msg,
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code":        rep.Code,
				"msg":         rep.Msg,
				"admin_token": token,
				"user_name":   rep.Username,
			})
		}

	}
}

func FrontUserList(ctx *gin.Context) {
	currentPage := ctx.DefaultQuery("currentPage", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")

	service := grpc.NewService()
	admin_user_service := user_srv.NewAdminUserService("xlf.user.srv.user_srv", service.Client())
	rep, err := admin_user_service.FrontUserList(context.TODO(), &user_srv.FrontUsersRequest{
		CurrentPage: utils.StrToInt(currentPage),
		PageSize:    utils.StrToInt(pageSize),
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "没有查询到数据",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":         200,
			"msg":          "成功",
			"front_users":  rep.FrontUsers,
			"total":        rep.Total,
			"current_page": rep.CurrentPage,
			"page_size":    rep.PageSize,
		})
	}

}
