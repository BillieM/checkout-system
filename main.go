package main

import (
	"github.com/billiem/checkout-system/checkout"
)

func main() {
	err := checkout.CheckoutCLI()
	if err != nil {
		panic(err)
	}
}
