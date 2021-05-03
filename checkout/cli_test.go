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
		expected string // expected string written to io.writer passed into CheckoutCli func
		expErr   bool
	}{
		{
			"1: example data",
			[]string{"./checkout_system", "-products=../testdata/product_sets/1.json", "../testdata/checkout_sets/1.json"},
			"checkout file: ../testdata/checkout_sets/1.json\nproducts file: ../testdata/product_sets/1.json\ntotal value of checkout: 284\n",
			false,
		},
		{
			"2: example checkout data with product data with no offers",
			[]string{"./checkout_system", "-products=../testdata/product_sets/2.json", "../testdata/checkout_sets/1.json"},
			"checkout file: ../testdata/checkout_sets/1.json\nproducts file: ../testdata/product_sets/2.json\ntotal value of checkout: 304\n",
			false,
		},
		{
			"3: checkout data with 0 quantities and example product data",
			[]string{"./checkout_system", "-products=../testdata/product_sets/1.json", "../testdata/checkout_sets/4.json"},
			"checkout file: ../testdata/checkout_sets/4.json\nproducts file: ../testdata/product_sets/1.json\ntotal value of checkout: 0\n",
			false,
		},
		{
			"4: checkout data with negative quantities and example product data",
			[]string{"./checkout_system", "-products=../testdata/product_sets/4.json", "../testdata/checkout_sets/7.json"},
			"",
			true,
		},
		{
			"5: example checkout data with product data with negative prices",
			[]string{"./checkout_system", "-products=../testdata/product_sets/5.json", "../testdata/checkout_sets/1.json"},
			"checkout file: ../testdata/checkout_sets/1.json\nproducts file: ../testdata/product_sets/5.json\ntotal value of checkout: -111087\n",
			false,
		},
		{
			"6: non-existent checkout file",
			[]string{"./checkout_system", "-products=../testdata/product_sets/1.json", "../testdata/checkout_sets/fake.json"},
			"",
			true,
		},
		{
			"7: non-existent product file",
			[]string{"./checkout_system", "-products=../testdata/product_sets/fake.json", "../testdata/checkout_sets/1.json"},
			"",
			true,
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
		{
			"2: only checkout arg given",
			[]string{"./checkout_system", "./other_checkout_data.json"},
			checkout.ArgInfo{
				"./other_checkout_data.json",
				"./product_data.json",
			},
		},
		{
			"3: only products flag given",
			[]string{"./checkout_system", "-products=./other_products_data.json"},
			checkout.ArgInfo{
				"./checkout_data.json",
				"./other_products_data.json",
			},
		},
		{
			"4: checkout arg/ products flag both given",
			[]string{"./checkout_system", "-products=./other_products_data.json", "./other_checkout_data.json"},
			checkout.ArgInfo{
				"./other_checkout_data.json",
				"./other_products_data.json",
			},
		},
		{
			"5: checkout arg/ products flag both given, + additional positional arg",
			[]string{"./checkout_system", "-products=./other_products_data.json", "./other_checkout_data.json", "./ignore_products_data.json"},
			checkout.ArgInfo{
				"./other_checkout_data.json",
				"./other_products_data.json",
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
