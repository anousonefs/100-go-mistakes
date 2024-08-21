package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	// bad error handling
	/* if _, err := GetRoute(90, 170, 91, -182); err != nil { */
	/* 	fmt.Printf("GetRoute(): %v\n", err) */
	/* } */

	// good
	/* if _, err := GetRoute2(90, 170, 91, -182); err != nil { */
	/* 	fmt.Printf("GetRoute2(): %v\n", err) */
	/* } */

	// best
	if _, err := GetRoute3(90, 170, 91, -182); err != nil {
		fmt.Printf("GetRoute3(): %v\n", err)
		fmt.Printf("GetRoute3(): %v\n", errors.Unwrap(err))
	}
}

type Route struct {
}

func GetRoute(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates(srcLat, srcLng)
	if err != nil {
		log.Println("failed to validate source coordinates")
		return Route{}, err
	}
	err = validateCoordinates(dstLat, dstLng)
	if err != nil {
		log.Println("failed to validate target coordinates")
		return Route{}, err
	}
	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func validateCoordinates(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		log.Printf("invalid latitude: %f", lat)
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng > 180.0 || lng < -180.0 {
		log.Printf("invalid longitude: %f", lng)
		return fmt.Errorf("invalid longitude: %f", lng)
	}
	return nil
}

func getRoute(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	_, _, _, _ = srcLat, srcLng, dstLat, dstLng
	return Route{}, nil
}

func validateCoordinates2(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng > 180.0 || lng < -180.0 {
		return fmt.Errorf("invalid longitude: %f", lng)
	}
	return nil
}

// source error
func GetRoute2(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates2(srcLat, srcLng)
	if err != nil {
		return Route{}, err
	}
	err = validateCoordinates2(dstLat, dstLng)
	if err != nil {
		return Route{}, err
	}
	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

// wraped error
func GetRoute3(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates2(srcLat, srcLng)
	if err != nil {
		return Route{}, fmt.Errorf("failed to validate source coordinates: %w", err)
	}
	err = validateCoordinates2(dstLat, dstLng)
	if err != nil {
		return Route{}, fmt.Errorf("failed to validate target coordinates: %w", err)
	}
	return getRoute(srcLat, srcLng, dstLat, dstLng)
}
