package checkout_test

import (
	"testing"

	"github.com/billiem/checkout-system/checkout"
)

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
			checkout.CheckoutLine{
				"A",
				3,
			},
			map[string]checkout.Product{"A": {
				Price:         50,
				OfferQuantity: 3,
				OfferPrice:    140,
			}},
			140,
			false,
		},
		{
			"2: example data prod 2",
			checkout.CheckoutLine{
				"B",
				3,
			},
			map[string]checkout.Product{"B": {
				Price:         35,
				OfferQuantity: 2,
				OfferPrice:    60,
			}},
			95,
			false,
		},
		{
			"3: example data prod 3",
			checkout.CheckoutLine{
				"C",
				1,
			},
			map[string]checkout.Product{"C": {
				Price: 25,
			}},
			25,
			false,
		},
		{
			"4: example data prod 4",
			checkout.CheckoutLine{
				"D",
				2,
			},
			map[string]checkout.Product{"D": {
				Price: 12,
			}},
			24,
			false,
		},
		{
			"5: negative checkout line quantity",
			checkout.CheckoutLine{
				"D",
				-3,
			},
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
			checkout.CheckoutLine{
				"D",
				1,
			},
			map[string]checkout.Product{
				"D": {
					Price:         15,
					OfferQuantity: -15,
					OfferPrice:    10,
				},
			},
			0,
			true,
		},
		{
			"7: checkout line product code not in products list",
			checkout.CheckoutLine{
				"A",
				3,
			},
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
