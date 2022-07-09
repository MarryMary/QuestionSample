package Session

import (
	"context"
	"github.com/elliotchance/phpserialize"
	"github.com/go-redis/redis/v9"
)

func Read(Key string) map[interface{}]interface{} {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var ctx = context.Background()

	Key = "PHPREDIS_SESSION:" + Key

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

func Get(Key string, Purpose any) any {
	if IsIn(Key, Purpose) {
		SESSION := Read(Key)
		return SESSION[Purpose]
	} else {
		return false
	}
}

func IsIn(Key string, Purpose any) bool {
	SESSION := Read(Key)

	if _, ok := SESSION[Purpose]; ok {
		return true
	} else {
		return false
	}
}
