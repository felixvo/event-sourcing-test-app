package main

import "github.com/go-redis/redis/v7"

func newRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-14450.c1.asia-northeast1-1.gce.cloud.redislabs.com:14450",
		Password: "37uaACndCvuQ1heADnHkishnAhMmosWq", // no password set
		DB:       0,                                  // use default DB
	})

	_, err := client.Ping().Result()
	return client, err

}
