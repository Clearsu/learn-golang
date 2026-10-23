package main

import (
	"fmt"
	"errors"
)

func findOrder(id string) (string, error) {
	if id == "order_1" {
		return "BoringBook", nil
	}
	return "", errors.New("Order not found.")
}

func main() {
	order, err := findOrder("order_3")

	if err != nil {
		fmt.Println("Failed: ", err)
		return
	}
	fmt.Println("Success: ", order)
}
