package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senhaDoToken")
	bookHotel(ctx, "token")
}

func bookHotel(ctx context.Context, key string) {
	token := ctx.Value(key).(string)
	fmt.Println(key, token)
}
