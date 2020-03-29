// package main

// import (
// 	"fmt"
// )

// type DesiredService struct {
// 	// The JSON tags are redundant here. See below.
// 	Name string `json:"Name"`
// 	Code int    `json:"Code"`
// }

// func main() {
// 	services := []DesiredService{
// 		DesiredService{"account-lists", 1},
// 		DesiredService{"access-control", 2},
// 		DesiredService{"activity-logs", 4},
// 		DesiredService{"taxes", 5},
// 	}

// 	for _, service := range services {
// 		// if service.Code == 8 {
// 		fmt.Println("\n", service.Name == "test")
// 		fmt.Println("\n", service.Code)
// 		// }
// 	}
// }
