package seckill

import (
	"fmt"
	"net/http"
	"web/rabbitmq"
	"web/utils"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

//// 秒杀接口
func SecKill(ctx *gin.Context) {
	id := ctx.PostForm("id")
	front_user_email, exist := ctx.Get("front_user_name")

	if !exist {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "token有问题",
		})
	}

	qe := rabbitmq.QueueAndExchange{
		QueueName:    "web.order_queue",
		ExchangeName: "web.order_exchange",
		ExchangeType: "direct",
		RoutingKey:   "web.order",
	}
	mq := rabbitmq.NewRabbitMq(qe)

	mq.ConnMq()
	mq.OpenChan()

	defer func() {
		mq.CloseMq()
	}()
	defer func() {
		mq.CloseChan()
	}()

	order_map := map[string]interface{}{
		"uemail": front_user_email,
		"pid":    id,
	}

	mq.PublishMsg(utils.MapToStr(order_map))

	////grpc 通信
	//service := grpc.NewService()
	//seckill_service := seckill_srv.NewSecKillService("xlf.seckill.srv.seckill_srv",service.Client())
	//rep,_ := seckill_service.FrontSecKill(context.TODO(),&seckill_srv.SecKillRequest{
	//	Id:utils.StrToInt(id),
	//	FrontUserEmail:front_user_email.(string),
	//})

	ctx.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  "下单中，请稍后",
	})

}

func GetSeckillResult(ctx *gin.Context) {

	uemail, exist := ctx.Get("front_user_name")

	fmt.Println("==============")
	fmt.Println(exist)
	if exist {
		conn, err := redis.Dial("tcp", "81.68.142.83:6379")
		if err != nil {
			fmt.Println("连接出错")
		}

		ret, err_r := redis.String(conn.Do("get", uemail))

		if err_r == nil {
			ret_map := utils.StrToMap(ret)
			fmt.Println(ret_map)
			ctx.JSON(http.StatusOK, gin.H{ // 说明从redis里面获取到了数据，
				"code": 200,
				"msg":  ret_map["msg"],
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
		})
		return
	}

}
