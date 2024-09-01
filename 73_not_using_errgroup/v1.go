package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Circle struct {
	// Circle fields here
	Radius int
}

type Result struct {
	// Result fields here
	Area float64
}

func foo(ctx context.Context, circle Circle) (Result, error) {
	// Simulate some processing and return a Result
	// Example: calculating the area of the circle
	if circle.Radius == 3 {
		return Result{}, errors.New("radius should not equal 3")
	}
	if circle.Radius == 7 {
		return Result{}, errors.New("radius should not equal 7")
	}
	time.Sleep(500 * time.Millisecond)
	area := 3.14 * float64(circle.Radius) * float64(circle.Radius)
	return Result{Area: area}, nil
}

func handler(ctx context.Context, circles []Circle) ([]Result, error) {
	println("handler 1")
	results := make([]Result, len(circles))
	wg := sync.WaitGroup{}
	wg.Add(len(results))
	errChan := make(chan error, len(circles)) // Channel to capture errors

	for i, circle := range circles {
		i := i
		circle := circle
		go func() {
			defer wg.Done()
			result, err := foo(ctx, circle)
			if err != nil {
				errChan <- err
				fmt.Printf("err: %v\n", circle.Radius)
				return
			}
			fmt.Printf("result %v: %v\n", circle.Radius, result)
			results[i] = result
		}()
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return nil, <-errChan
	}

	return results, nil
}
