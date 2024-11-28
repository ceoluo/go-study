package main

import (
	"fmt"
	"log"
)

type Route struct{}

// GetRoute1
// example:
// func main() {
// 	srcLat, srcLng, dstLat, dstLng := float32(91), float32(91), float32(91), float32(91)
// 	_, err := GetRoute1(srcLat, srcLng, dstLat, dstLng)
// 	if err != nil {
// 		fmt.Printf("%v\n", err)
// 	}
// }
// output:
// 2024/01/24 11:54:37 invalid latitude: 91.000000
// 2024/01/24 11:54:37 failed to validate source coordinates
// invalid latitude: 91.000000
func GetRoute1(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates1(srcLat, srcLng)
	if err != nil {
		log.Println("failed to validate source coordinates")
		return Route{}, err
	}

	err = validateCoordinates1(dstLat, dstLng)
	if err != nil {
		log.Println("failed to validate target coordinates")
		return Route{}, err
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func validateCoordinates1(lat, lng float32) error {
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

// GetRoute2
// example:
// func main() {
// 	srcLat, srcLng, dstLat, dstLng := float32(91), float32(91), float32(91), float32(91)
// 	_, err := GetRoute2(srcLat, srcLng, dstLat, dstLng)
// 	if err != nil {
// 		fmt.Printf("%v\n", err)
// 	}
// }
// output:
// invalid latitude: 91.000000
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

func validateCoordinates2(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng > 180.0 || lng < -180.0 {
		return fmt.Errorf("invalid longitude: %f", lng)
	}
	return nil
}

// GetRoute3
// example:
// func main() {
// 	srcLat, srcLng, dstLat, dstLng := float32(91), float32(91), float32(91), float32(91)
// 	_, err := GetRoute3(srcLat, srcLng, dstLat, dstLng)
// 	if err != nil {
// 		fmt.Printf("%v\n", err)
// 	}
// }
// output:
// failed to validate source coordinates: invalid latitude: 91.000000
func GetRoute3(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates2(srcLat, srcLng)
	if err != nil {
		return Route{},
			fmt.Errorf("failed to validate source coordinates: %w", err)
	}

	err = validateCoordinates2(dstLat, dstLng)
	if err != nil {
		return Route{},
			fmt.Errorf("failed to validate target coordinates: %w", err)
	}

	return getRoute(srcLat, srcLng, dstLat, dstLng)
}

func getRoute(lat, lng, lat2, lng2 float32) (Route, error) {
	return Route{}, nil
}

func main() {
	srcLat, srcLng, dstLat, dstLng := float32(91), float32(91), float32(91), float32(91)
	_, err := GetRoute3(srcLat, srcLng, dstLat, dstLng)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
