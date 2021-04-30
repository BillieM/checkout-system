package main

import (
	"fmt"

	"github.com/billiem/checkout-system/checkout"
)

// Takes filename as an argument & prints parsed value to stdout
func main() {
	fmt.Println("hello world")
	checkout.DecodeCheckoutData("./testdata/checkout_sets/1.json")
}
