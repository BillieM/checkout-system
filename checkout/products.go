/*
Package checkout provides utilities to calculate a total checkout price from JSON data.
*/
package checkout

import (
	"errors"
)

type (
	// CheckoutLine stores information about a particular line parsed from checkout data
	//
	// Contains a product code (e.g. "A") represented as a string, and an integer Quantity of the item (e.g. 5). A negative Quantity is invalid
	//
	// DecodeCheckoutData (io.go) returns data from checkout JSON as a slice of CheckoutLine
	CheckoutLine struct {
		Code     string
		Quantity int
	}

	// Product stores price information about a particular product.
	//
	// Price is an integer value (e.g. 6), acting as the unit cost, negative prices are allowed to allow discounting functionality.
	// OfferQuantity is the quantity which must be purchased to benefit from the offer,
	// with OfferPrice being the price of the given OfferQuantity (e.g. if OfferQuantity is 3, and OfferPrice is 150, 3 items will cost 150).
	// If OfferQuantity is 0/ not given, offers will be ignored. A negative OfferQuantity is invalid
	//
	// DecodePriceData (io.go) returns a map of [string: Product Code]Product
	Product struct {
		Price         int
		OfferQuantity int
		OfferPrice    int
	}
)

// ProcessCheckout is a function from calculating the value of a checkout.
//
// It accepts the path to the checkout json file, and the path to the products list json file.
//
// Returned is the total value from GetCheckoutPrice and any errors that have occured calling other functions.
func ProcessCheckout(checkoutPath string, productsPath string) (int, error) {

	// get checkout line arr and products map
	checkoutLines, err := DecodeCheckoutData(checkoutPath)
	if err != nil {
		return 0, err
	}
	products, err := DecodeProductData(productsPath)
	if err != nil {
		return 0, err
	}

	checkoutPrice, err := GetCheckoutPrice(checkoutLines, products)
	if err != nil {
		return 0, err
	}

	return checkoutPrice, nil
}

// GetCheckoutLinePrice is a method for CheckoutLine which also accepts a map of [productCode]Product representing product prices and their current offers.
//
// returns an error if the checkout line quantity is negative, or if the offer quantity is negative
func (cL CheckoutLine) GetCheckoutLinePrice(products map[string]Product) (int, error) {

	lineTotal := 0

	// check for invalid checkout quantity.
	if cL.Quantity < 0 {
		return 0, errors.New("checkout line quantity cannot be negative")
	}

	// check prod key in products map, if not return err else continue
	if prod, ok := products[cL.Code]; ok {
		// check for invalid offer quantity
		if prod.OfferQuantity < 0 {
			return 0, errors.New("offer quantity cannot be negative")
		}
		// check if there is an offer to be used
		if prod.OfferQuantity > 0 {
			// offer exists, apply offer for as many items as possible, and normal price for the rest
			lineTotal += (cL.Quantity / prod.OfferQuantity) * prod.OfferPrice
			lineTotal += (cL.Quantity % prod.OfferQuantity) * prod.Price
		} else {
			// no offer, apply normal price for all items
			lineTotal += cL.Quantity * prod.Price
		}
	} else {
		return 0, errors.New("no product code or product code not found in products map")
	}

	return lineTotal, nil
}

// GetCheckoutPrice accepts a slice of CheckoutLine and a map of [productCode]Product representing product prices.
//
// Checkout lines are iterated over, for each the CheckoutLine method is called, and the lineTotal is added to the sum
//
// If an error occurs in GetCheckoutLinePrice, it is returned from this function.
func GetCheckoutPrice(cLSlice []CheckoutLine, products map[string]Product) (int, error) {

	checkoutTotal := 0

	// loop over checkout lines, get their price and add it to the total checkout price
	for _, cL := range cLSlice {

		cLPrice, err := cL.GetCheckoutLinePrice(products)
		if err != nil {
			return 0, err
		}

		checkoutTotal += cLPrice
	}

	return checkoutTotal, nil
}
