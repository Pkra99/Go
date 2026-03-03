package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    streamKey := "events"
    args := &redis.XAddArgs{
        Stream: streamKey,
        MaxLen: 1000,
        Approx: true,
        Values: map[string]interface{}{
            "user":   "maneshwar",
            "action": "upload",
            "time":   time.Now().Format(time.RFC3339),
        },
    }
    id, err := rdb.XAdd(ctx, args).Result()
    if err != nil {
        log.Fatalf("XAdd failed: %v", err)
    }

    fmt.Printf("Written to stream with ID: %s\n", id)
}