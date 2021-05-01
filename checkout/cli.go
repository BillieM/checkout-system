package checkout

import (
	"flag"
	"fmt"
	"os"
)

const (
	// Default checkout data filePath
	CheckoutPath = "./checkout_data.json"

	// Default price data filePath
	PricesPath = "./prices_config.json"
)

//ArgInfo is returned from getArgInfo and contains the filepaths for the checkout file/ prices file (if they are given)
//
//Filepaths may be relative or absolute
type ArgInfo struct {
	checkoutFile string // checkout json file path
	pricesFile   string // prices json file path
}

// getArgInfo returns an instance of ArgInfo.
//
// if the checkout info file path has not been given, or the prices flag has not been given,
//checkoutFile/ pricesFile will be returned as "" respectively.
//
//Filepaths may be relative or absolute.
func getArgInfo() ArgInfo {

	// get prices flag value for prices file
	pricesPath := flag.String("prices", "", "optional filepath to prices JSON")
	flag.Parse()

	// get first positional argument for checkout file
	checkoutPath := flag.Arg(0)

	// set to default paths if flag empty/ checkout argument not given

	if *pricesPath == "" {
		*pricesPath = PricesPath
	}

	if checkoutPath == "" {
		checkoutPath = CheckoutPath
	}

	return ArgInfo{
		checkoutFile: checkoutPath,
		pricesFile:   *pricesPath,
	}
}

// Called from main package, primary entry point.
//
// Program takes a filename as an argument, expecting a json file of checkout lines,
// If no filename argument is given, the default checkout dataset will instead be used.
//
// An optional prices flag can also be given to specify a path to a different price list.
func CheckoutCLI() error {
	// --help info
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <inJSONLocation>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	argInfo := getArgInfo()

	// logic to extract from json/ calc checkout value
	result, err := ProcessCheckout(argInfo.checkoutFile, argInfo.pricesFile)

	if err != nil {
		return err
	}

	fmt.Printf("checkout file: %s\nprices file: %s\ntotal value of checkout: %v\n", argInfo.checkoutFile, argInfo.pricesFile, result)

	return nil
}
