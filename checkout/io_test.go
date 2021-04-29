package checkout_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/billiem/checkout-system/checkout"
)

func Test_DecodeCheckoutData(t *testing.T) {

	/*
		tests the DecodeCheckoutData function
	*/

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
				{"A", 60},
				{"B", 23},
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
	}
	for _, testCase := range testCases {
		// run subtest for each testCase
		t.Run(testCase.name, func(t *testing.T) {
			testDataPath := "../testdata/checkout_sets/" + testCase.fileName
			out, err := checkout.DecodeCheckoutData(testDataPath)
			// check for err when decoding file
			if (err != nil) != testCase.expErr {
				t.Errorf("case: %s, expected err: %v, got err: %v", testCase.name, testCase.expErr, err)
			}
			// check returned []CheckoutLine equal to expected
			if !reflect.DeepEqual(out, testCase.expected) {
				fmt.Println(out, testCase.expected)
				t.Errorf("case: %s, []CheckoutLine different from expected", testCase.name)
			}
		})
	}
}
