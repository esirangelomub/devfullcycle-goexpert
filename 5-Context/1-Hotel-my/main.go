package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	//ctx, cancel := context.WithCancel(ctx)
	ctx, cancel := context.WithTimeout(ctx, 7*time.Second)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <-time.After(6 * time.Second):
		fmt.Println("Hotel booking successful.")
		return
	}
}
