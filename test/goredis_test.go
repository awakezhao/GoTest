package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestConnnection(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	fmt.Println(rdb.Options().Addr)
	fmt.Println(rdb.Options().Network)
}

func TestOp(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()

	rdb.Set(ctx, "A:B:C:D:key", 100, 0)

	val, err := rdb.Get(ctx, "A:B:C:D:key").Result()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	get := rdb.Get(ctx, "key")
	fmt.Println(get.Val(), get.Err())

	var min = "(" + "0"
	ret := rdb.ZRangeByScoreWithScores(context.Background(), "s15:timestamp_rank", &redis.ZRangeBy{
		Min:    min,
		Max:    "+inf",
		Offset: 0,
		Count:  int64(10),
	})
	for _, val := range ret.Val() {
		fmt.Println(val.Member)
	}
}
