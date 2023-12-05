package helper

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	SKUPrefix = "SKU"
)

func GenerateSKU() string {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 10000000 and 90000000
	randomNumber := rand.Intn(90000000) + 10000000

	// Combine components to create the SKU
	sku := fmt.Sprintf("%s-%d", SKUPrefix, randomNumber)

	return sku
}
