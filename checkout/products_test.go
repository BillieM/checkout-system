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
		{
			"13: unused offer",
			checkout.CheckoutLine{"D", 4},
			map[string]checkout.Product{
				"D": {
					Price: 12, OfferQuantity: 100, OfferPrice: 4,
				},
			},
			48,
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
		{
			"4: no checkout lines",
			[]checkout.CheckoutLine{},
			map[string]checkout.Product{},
			0,
			false,
		},
		{
			"5: products map missing a product",
			[]checkout.CheckoutLine{
				{"A", 66},
				{"B", 3123},
				{"C", 661},
				{"D", 4},
			},
			map[string]checkout.Product{
				"A": {Price: 50, OfferQuantity: 3, OfferPrice: 140},
				"B": {Price: 35, OfferQuantity: 2, OfferPrice: 60},
				"C": {Price: 25},
			},
			0,
			true,
		},
		{
			"6: no offers",
			[]checkout.CheckoutLine{
				{"A", 12},
				{"B", 4},
				{"C", 2},
				{"D", 1},
			},
			map[string]checkout.Product{
				"A": {Price: 50},
				"B": {Price: 35},
				"C": {Price: 25},
				"D": {Price: 14},
			},
			804,
			false,
		},
		{
			"7: all products have offers",
			[]checkout.CheckoutLine{
				{"A", 12},
				{"B", 4},
				{"C", 22},
			},
			map[string]checkout.Product{
				"A": {Price: 50, OfferQuantity: 4, OfferPrice: 150},
				"B": {Price: 35, OfferQuantity: 7, OfferPrice: 100},
				"C": {Price: 25, OfferQuantity: 2, OfferPrice: 28},
			},
			898,
			false,
		},
		{
			"8: negative prices and negative offer prices",
			[]checkout.CheckoutLine{
				{"A", 12},
				{"B", 4},
				{"C", 15},
			},
			map[string]checkout.Product{
				"A": {Price: -50},
				"B": {Price: -12, OfferQuantity: 5, OfferPrice: -30},
				"C": {Price: -4, OfferQuantity: 6, OfferPrice: -20},
			},
			-700,
			false,
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

// Test_ProcessCheckout tests the ProcessCheckout function.
//
// It passes in a filepath to the checkout json data, and a filepath to the products json data,
// testing for the correct checkout value and whether or not an error was expected.
func Test_ProcessCheckout(t *testing.T) {
	testCases := []struct {
		name         string
		checkoutFile string
		productsFile string
		expected     int
		expErr       bool
	}{
		{
			"1: given example data",
			"1.json",
			"1.json",
			284,
			false,
		},
		{
			"2: example checkout data with product data with no offers",
			"1.json",
			"2.json",
			304,
			false,
		},
		{
			"3: checkout data with 0 quantities and example product data",
			"4.json",
			"1.json",
			0,
			false,
		},
		{
			"4: checkout data with negative quantities and example product data",
			"7.json",
			"1.json",
			0,
			true,
		},
		{
			"5: example checkout data with product data with negative prices",
			"1.json",
			"5.json",
			-111087,
			false,
		},
		{
			"6: non-existent checkout file",
			"fake.json",
			"1.json",
			0,
			true,
		},
		{
			"7: non-existent products file",
			"1.json",
			"fake.json",
			0,
			true,
		},
	}

	// loop over test cases
	for _, testCase := range testCases {
		// run subtest for each test case
		t.Run(testCase.name, func(t *testing.T) {
			checkoutPath := "../testdata/checkout_sets/" + testCase.checkoutFile
			productsPath := "../testdata/product_sets/" + testCase.productsFile
			result, err := checkout.ProcessCheckout(checkoutPath, productsPath)
			// check if err expected
			if (err != nil) != testCase.expErr {
				t.Errorf("expected error: %v, got err: %s", testCase.expErr, err)
			}
			// compare result val to expected val
			if result != testCase.expected {
				t.Errorf("expected checkout value of: %v, got checkout value of %v", testCase.expected, result)
			}
		})
	}
}
