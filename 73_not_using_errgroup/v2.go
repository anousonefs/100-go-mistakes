package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func handler2(ctx context.Context, circles []Circle) ([]Result, error) {
	println("handler 2")
	results := make([]Result, len(circles))
	g, ctx := errgroup.WithContext(ctx)
	for i, circle := range circles {
		i := i
		circle := circle
		g.Go(func() error {
			result, err := foo(ctx, circle)
			if err != nil {
				fmt.Printf("err: %v\n", circle.Radius)
				return err
			}
			results[i] = result
			fmt.Printf("result %v: %v\n", circle.Radius, result)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return results, nil
}
