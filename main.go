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

/*

to do

ProcessCheckout tests

cli tests

main test ?? (maybe ??)

readme

*/
