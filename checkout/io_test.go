package checkout_test

import (
	"reflect"
	"testing"

	"github.com/billiem/checkout-system/checkout"
)

// Tests the DecodeCheckoutData function using data from testdata/checkout_sets
func Test_DecodeCheckoutData(t *testing.T) {

	testCases := []struct {
		name     string
		fileName string
		expected []checkout.CheckoutLine
		expErr   bool
	}{
		{
			"1: given example case",
			"1.json",
			[]checkout.CheckoutLine{
				{"A", 3},
				{"B", 3},
				{"C", 1},
				{"D", 2},
			},
			false,
		},
		{
			"2: normal data",
			"2.json",
			[]checkout.CheckoutLine{
				{"A", 4},
				{"B", 1},
				{"C", 2},
				{"D", 6},
			},
			false,
		},
		{
			"3: only 3 lines",
			"3.json",
			[]checkout.CheckoutLine{
				{"A", 4},
				{"B", 0},
				{"D", 2},
			},
			false,
		},
		{
			"4: all 0 quantity",
			"4.json",
			[]checkout.CheckoutLine{
				{"A", 0},
				{"B", 0},
				{"C", 0},
				{"D", 0},
			},
			false,
		},
		{
			"5: only 3 lines, higher vals, 1 qty 0",
			"5.json",
			[]checkout.CheckoutLine{
				{"A", 94124},
				{"B", 999991023},
				{"C", 0},
			},
			false,
		},
		{
			"6: 1 line with negative qty",
			"6.json",
			[]checkout.CheckoutLine{
				{"A", 60},
				{"B", 23},
				{"C", 0},
				{"D", -1},
			},
			false,
		},
		{
			"7: 4 lines with negative qty",
			"7.json",
			[]checkout.CheckoutLine{
				{"A", -4},
				{"B", -7},
				{"C", -9},
				{"D", -2},
			},
			false,
		},
		{
			"8: blank txt file",
			"0.txt",
			[]checkout.CheckoutLine{},
			true,
		},
		{
			"9: non-existent file",
			"fake.json",
			[]checkout.CheckoutLine{},
			true,
		},
		{
			"10: invalid json format",
			"8.json",
			[]checkout.CheckoutLine{},
			true,
		},
	}
	for _, testCase := range testCases {
		// run subtest for each testCase
		t.Run(testCase.name, func(t *testing.T) {
			testDataPath := "../testdata/checkout_sets/" + testCase.fileName
			result, err := checkout.DecodeCheckoutData(testDataPath)
			// check if err expected
			if (err != nil) != testCase.expErr {
				t.Errorf("case: %s, expected err: %v, got err: %v", testCase.name, testCase.expErr, err)
			}
			// check returned []CheckoutLine equal to expected
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("case: %s, []CheckoutLine different from expected", testCase.name)
			}
		})
	}
}

// Tests the DecodeProductData function using data from testdata/product_sets
func Test_DecodeProductData(t *testing.T) {

	testCases := []struct {
		name     string
		fileName string
		expected map[string]checkout.Product
		expErr   bool
	}{
		{
			"1: given example testData",
			"1.json",
			map[string]checkout.Product{
				"A": {
					Price:         50,
					OfferQuantity: 3,
					OfferPrice:    140,
				},
				"B": {
					Price:         35,
					OfferQuantity: 2,
					OfferPrice:    60,
				},
				"C": {
					Price: 25,
				},
				"D": {
					Price: 12,
				},
			},
			false,
		},
		{
			"2: valid input data with no offers",
			"2.json",
			map[string]checkout.Product{
				"A": {
					Price: 50,
				},
				"B": {
					Price: 35,
				},
				"C": {
					Price: 25,
				},
				"D": {
					Price: 12,
				},
			},
			false,
		},

		{
			"3: only 3 lines",
			"3.json",
			map[string]checkout.Product{
				"A": {
					Price: 50,
				},
				"B": {
					Price: 35,
				},
				"C": {
					Price: 25,
				},
			},
			false,
		},
		{
			"4: all price 0",
			"4.json",
			map[string]checkout.Product{
				"A": {
					Price: 0,
				},
				"B": {
					Price: 0,
				},
				"C": {
					Price: 0,
				},
				"D": {
					Price: 0,
				},
			},
			false,
		},
		{
			"5: all negative prices",
			"5.json",
			map[string]checkout.Product{
				"A": {
					Price: -412,
				},
				"B": {
					Price: -6123,
				},
				"C": {
					Price: -91234,
				},
				"D": {
					Price:         -124,
					OfferQuantity: 5,
					OfferPrice:    -555,
				},
			},
			false,
		},
	}

	for _, testCase := range testCases {
		// run subtest for each test case
		t.Run(testCase.name, func(t *testing.T) {
			testDataPath := "../testdata/product_sets/" + testCase.fileName
			result, err := checkout.DecodeProductData(testDataPath)
			// check if err expected
			if (err != nil) != testCase.expErr {
				t.Errorf("case: %s, expected err: %v, got err: %v", testCase.name, testCase.expErr, err)
			}
			// check returned map[string]Product equal to expected
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("case: %s, map[string]Product different from expected", testCase.name)
			}
		})
	}
}
