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

ProcessCheckout logic and tests

GetCheckoutPrice tests

cli tests

main test ?? (maybe ??)

*/
