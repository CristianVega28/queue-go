package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	RedisSrvc struct{}
)

func (r *RedisSrvc) Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}

func (r *RedisSrvc) Get(ctx context.Context, key string, valueDefault any) any {
	rdb := r.Connect()
	defer rdb.Close()

	if !r.Has(ctx, key) {
		return valueDefault
	}

	t, err := rdb.Type(ctx, key).Result()

	if err != nil {
		fmt.Println(err.Error())
	}

	var result any
	switch t {
	case "string":
		result_rdb, _ := rdb.Get(ctx, key).Result()
		var unmarshaledResult any
		err := json.Unmarshal([]byte(result_rdb), &unmarshaledResult)
		if err != nil {
			fmt.Println(err.Error())
		}
		result = unmarshaledResult
	case "hash":
		result, _ = rdb.HGetAll(ctx, key).Result()
	case "list":

	}

	return result
}

func (r *RedisSrvc) Has(ctx context.Context, key string) bool {
	rdb := r.Connect()
	defer rdb.Close()

	exist, err := rdb.Exists(ctx, key).Result()

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if exist == 1 {
		return true
	} else {
		return false
	}
}

func (r *RedisSrvc) HSet(ctx context.Context, key string, values any, ttl time.Duration) error {
	rdb := r.Connect()
	defer rdb.Close()

	err := rdb.HSet(ctx, key, values).Err()

	if err != nil {
		return err
	}

	if ttl != 0 {
		err := rdb.Expire(ctx, key, ttl).Err()

		if err != nil {
			return err
		}
	}

	return nil

}

func (r *RedisSrvc) Set(ctx context.Context, key string, values any, ttl time.Duration) error {
	rdb := r.Connect()
	defer rdb.Close()

	err := rdb.Set(ctx, key, values, ttl).Err()

	if err != nil {
		return err
	}

	return nil

}
