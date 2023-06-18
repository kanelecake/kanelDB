package main

import (
	"context"
	"fmt"
	kaneldb "github.com/kanelecake/kanelDB"
	"time"
)

var ctx = context.Background()

func main() {
	// create a database instance
	db := kaneldb.NewClient()

	// set to database
	ttl := 3 * time.Minute
	err := db.Set(ctx, "hello_key", "hello_world", &ttl).Err()
	if err != nil {
		panic(err)
	}

	// get a value
	result, getErr := db.Get(ctx, "hello_key").Result()
	if getErr != nil {
		panic(err)
	}

	fmt.Printf("%s", result)
}
