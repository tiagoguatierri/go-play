package main

import (
	"fmt"

	"github.com/tiagoguatierri/go-play/internal/core/entity"
)

func main() {
	customer := new(entity.Customer)

	customer.ID = "1"

	fmt.Printf("Customer ID: %v\n", customer)
}
