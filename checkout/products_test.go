package checkout_test

import (
	"testing"

	"github.com/billiem/checkout-system/checkout"
)

// Test_GetCheckoutLinePrice tests the GetCheckoutLinePrice function.
//
// It passes a map of [productCode]Product and a CheckoutLine to the function,
// checking for the correct expected line price, and whether or not an error is expected.
func Test_GetCheckoutLinePrice(t *testing.T) {
	testCases := []struct {
		name         string
		checkoutLine checkout.CheckoutLine
		products     map[string]checkout.Product
		expected     int
		expErr       bool
	}{
		{
			"1: example data prod 1",
			checkout.CheckoutLine{"A", 3},
			map[string]checkout.Product{"A": {
				Price: 50, OfferQuantity: 3, OfferPrice: 140}},
			140,
			false,
		},
		{
			"2: example data prod 2",
			checkout.CheckoutLine{"B", 3},
			map[string]checkout.Product{"B": {
				Price: 35, OfferQuantity: 2, OfferPrice: 60,
			}},
			95,
			false,
		},
		{
			"3: example data prod 3",
			checkout.CheckoutLine{"C", 1},
			map[string]checkout.Product{"C": {
				Price: 25,
			}},
			25,
			false,
		},
		{
			"4: example data prod 4",
			checkout.CheckoutLine{"D", 2},
			map[string]checkout.Product{"D": {
				Price: 12,
			}},
			24,
			false,
		},
		{
			"5: negative checkout line quantity",
			checkout.CheckoutLine{"D", -3},
			map[string]checkout.Product{
				"D": {
					Price: 10,
				},
			},
			0,
			true,
		},
		{
			"6: negative offer line quantity",
			checkout.CheckoutLine{"D", 1},
			map[string]checkout.Product{
				"D": {
					Price: 15, OfferQuantity: -15, OfferPrice: 10,
				},
			},
			0,
			true,
		},
		{
			"7: checkout line product code not in products list",
			checkout.CheckoutLine{"A", 3},
			map[string]checkout.Product{
				"B": {
					Price: 10,
				},
			},
			0,
			true,
		},
		{
			"8: empty checkout line",
			checkout.CheckoutLine{},
			map[string]checkout.Product{
				"A": {
					Price: 5,
				},
			},
			0,
			true,
		},
		{
			"9: negative price positive offer price",
			checkout.CheckoutLine{"A", 14},
			map[string]checkout.Product{
				"A": {
					Price: -44, OfferQuantity: 4, OfferPrice: 19,
				},
			},
			-31,
			false,
		},
		{
			"10: positive price negative offer price",
			checkout.CheckoutLine{"C", 25},
			map[string]checkout.Product{
				"C": {
					Price: 12, OfferQuantity: 3, OfferPrice: -2,
				},
			},
			-4,
			false,
		},
		{
			"11: negative price and negative offer price",
			checkout.CheckoutLine{"D", 66},
			map[string]checkout.Product{
				"D": {
					Price: -4, OfferQuantity: 3, OfferPrice: -5,
				},
			},
			-110,
			false,
		},
		{
			"12: large checkout line quantity",
			checkout.CheckoutLine{"F", 29124908},
			map[string]checkout.Product{
				"F": {
					Price: 4, OfferQuantity: 44, OfferPrice: 140,
				},
			},
			92670188,
			false,
		},
	}

	// loop over and run testcases
	for _, testCase := range testCases {
		// run subtest for each testCase
		t.Run(testCase.name, func(t *testing.T) {
			result, err := testCase.checkoutLine.GetCheckoutLinePrice(testCase.products)
			// check if err expected
			if (err != nil) != testCase.expErr {
				t.Errorf("expected error: %v, got err: %s", testCase.expErr, err)
			}
			// compare result val to expected val
			if result != testCase.expected {
				t.Errorf("expected line price of: %v, got line price of: %v", testCase.expected, result)
			}
		})
	}
}

// Test_GetCheckoutPrice tests the GetCheckoutPrice function.
//
// It passes a map of [productCode]Product and a slice of CheckoutLine to the function,
// checking for the correct expected checkout price, and whether or not an error is expected.
func Test_GetCheckoutPrice(t *testing.T) {
	testCases := []struct {
		name          string
		checkoutLines []checkout.CheckoutLine
		products      map[string]checkout.Product
		expected      int
		expErr        bool
	}{
		{
			"1: given example checkout data & given example products list",
			[]checkout.CheckoutLine{
				{"A", 3},
				{"B", 3},
				{"C", 1},
				{"D", 2},
			},
			map[string]checkout.Product{
				"A": {Price: 50, OfferQuantity: 3, OfferPrice: 140},
				"B": {Price: 35, OfferQuantity: 2, OfferPrice: 60},
				"C": {Price: 25},
				"D": {Price: 12},
			},
			284,
			false,
		},
		{
			"2: normal checkout data & given example products list",
			[]checkout.CheckoutLine{
				{"A", 66},
				{"B", 3123},
				{"C", 661},
				{"D", 21},
			},
			map[string]checkout.Product{
				"A": {Price: 50, OfferQuantity: 3, OfferPrice: 140},
				"B": {Price: 35, OfferQuantity: 2, OfferPrice: 60},
				"C": {Price: 25},
				"D": {Price: 12},
			},
			113552,
			false,
		},
		{
			"3: negative checkout quantity line & given example products list",
			[]checkout.CheckoutLine{
				{"A", 66},
				{"B", 3123},
				{"C", 661},
				{"D", -1},
			},
			map[string]checkout.Product{
				"A": {Price: 50, OfferQuantity: 3, OfferPrice: 140},
				"B": {Price: 35, OfferQuantity: 2, OfferPrice: 60},
				"C": {Price: 25},
				"D": {Price: 12},
			},
			0,
			true,
		},
	}

	// loop over and run test cases
	for _, testCase := range testCases {
		// run subtest for each testCase
		t.Run(testCase.name, func(t *testing.T) {
			result, err := checkout.GetCheckoutPrice(testCase.checkoutLines, testCase.products)
			// check if err expected
			if (err != nil) != testCase.expErr {
				t.Errorf("expected error: %v, got err: %s", testCase.expErr, err)
			}
			// compare result val to expected val
			if result != testCase.expected {
				t.Errorf("expected line price of: %v, got checkout price of: %v", testCase.expected, result)
			}
		})
	}
}
