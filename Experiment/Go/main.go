package main

import (
	"context"

	"github.com/elliotchance/phpserialize"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

func RedisSample(Key string) map[interface{}]interface{} {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	val2, err := rdb.Get(ctx, Key).Result()
	if err == redis.Nil {
		in := make(map[interface{}]interface{})
		return in
	} else if err != nil {
		panic(err)
	} else {
		in, err := phpserialize.UnmarshalAssociativeArray([]byte(val2))

		if err != nil {
			panic(err)
		}
		return in

	}
}

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		php_sess_id, err := c.Cookie("PHPSESSID")

		if err != nil {
			panic(err)
		}

		redis_sess_key := "PHPREDIS_SESSION:" + php_sess_id

		RedisSample(redis_sess_key)
	})
	engine.Run(":3000")
}
