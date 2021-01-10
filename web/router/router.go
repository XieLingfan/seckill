package all_router

import (
	"web/controller/product"
	"web/controller/seckill"
	"web/controller/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	user_group := router.Group("/user")
	product_group := router.Group("/product")
	seckill_group := router.Group("/seckill")

	user.Router(user_group)
	product.Router(product_group)
	seckill.Router(seckill_group)
}
