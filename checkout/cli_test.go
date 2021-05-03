package checkout_test

import (
	"bytes"
	"os"
	"reflect"
	"testing"

	"github.com/billiem/checkout-system/checkout"
)

func Test_CheckoutCLI(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expected string
		expErr   bool
	}{
		{
			"1: example data",
			[]string{"./checkout_system", "-products=../testdata/product_sets/1.json", "../testdata/checkout_sets/1.json"},
			"checkout file: ../testdata/checkout_sets/1.json\nproducts file: ../testdata/product_sets/1.json\ntotal value of checkout: 284\n",
			false,
		},
	}

	// loop over test cases
	for _, testCase := range testCases {
		// subtest for each test
		t.Run(testCase.name, func(t *testing.T) {

			// set args, io.writer & call CheckoutCLI function
			os.Args = testCase.args
			out := bytes.NewBuffer(nil)
			err := checkout.CheckoutCLI(out)

			// check if err expected
			if (err != nil) != testCase.expErr {
				t.Errorf("expected error: %v, got err: %s", testCase.expErr, err)
			}

			// check result
			if outStr := out.String(); outStr != testCase.expected {
				t.Errorf("expected:\n%sgot:\n%s", testCase.expected, outStr)
			}
		})
	}
}

func Test_GetArgInfo(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expected checkout.ArgInfo
	}{
		{
			"1: no arg/ flag given",
			[]string{"./checkout_system"},
			checkout.ArgInfo{
				"./checkout_data.json",
				"./product_data.json",
			},
		},
	}

	// loop over test cases
	for _, testCase := range testCases {
		// run subtest for each testcase
		t.Run(testCase.name, func(t *testing.T) {
			// set args and call GetArgInfo
			os.Args = testCase.args

			// check expected argInfo matches returned argInfo
			result := checkout.GetArgInfo()
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected: %v, got %v", testCase.expected, result)
			}
		})
	}

}
