package Session

import (
	"context"
	"github.com/elliotchance/phpserialize"
	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

func RedisSample(Key string) map[interface{}]interface{} {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	PHP_SESSION, err := rdb.Get(ctx, Key).Result()
	if err == redis.Nil {
		in := make(map[interface{}]interface{})
		return in
	} else if err != nil {
		panic(err)
	} else {
		in, err := phpserialize.UnmarshalAssociativeArray([]byte(PHP_SESSION))

		if err != nil {
			panic(err)
		}
		return in

	}
}
