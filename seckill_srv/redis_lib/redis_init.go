package redis_lib

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var Conn redis.Conn
var Err error

func init() {

	Conn, Err = redis.Dial("tcp", "81.68.142.83:6379")
	if Err != nil {
		fmt.Println(Err)
		panic(Err)
	}

}
