package main

import (
	"fmt"

	"github.com/billiem/checkout-system/checkout"
)

// cli tool that takes filename as an argument

func main() {
	fmt.Println("hello world")
	checkout.DecodeCheckoutData("./testdata/checkout_sets/1.json")
}
