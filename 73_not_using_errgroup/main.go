package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	circles := []Circle{
		{Radius: 2},
		{Radius: 3},
		{Radius: 4},
		{Radius: 5},
		{Radius: 7},
	}

	/* results, err := handler(ctx, circles) */
	results, err := handler2(ctx, circles)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i, result := range results {
		fmt.Printf("Result %d: Area = %f\n", i, result.Area)
	}
}
