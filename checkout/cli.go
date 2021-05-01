package checkout

import (
	"flag"
	"fmt"
	"os"
)

/*
ArgInfo is returned from getArgInfo and contains the filepaths for the checkout file/ prices file (if they are given)
Filepaths may be relative or absolute
*/
type ArgInfo struct {
	checkoutFile string // checkout json file path
	pricesFile   string // prices json file path
}

/*
getArgInfo returns an instance of ArgInfo.
if the checkout info file path has not been given, or the prices flag has not been given,
checkoutFile/ pricesFile will be returned as "" respectively.
Filepaths may be relative or absolute.
*/
func getArgInfo() ArgInfo {

	// get prices flag value for prices file
	pricesFile := flag.String("prices", "", "optional filepath to prices JSON")
	flag.Parse()

	// get first positional argument for checkout file
	checkoutFile := flag.Arg(0)

	// set to default paths if flag empty/ checkout argument not given

	if *pricesFile == "" {
		*pricesFile = "./prices_config.json"
	}

	if checkoutFile == "" {
		checkoutFile = "./checkout_data.json"
	}

	return ArgInfo{
		checkoutFile: checkoutFile,
		pricesFile:   *pricesFile,
	}
}

/*
called from main package, primary entry point.
Program takes a filename as an argument, expecting a json file of checkout lines,
If no filename argument is given, the default dataset will instead be used.
An optional prices flag can also be given to specify a path to a different price list.
*/
func CLI() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <inJSONLocation>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	argInfo := getArgInfo()

	result, err := ProcessCheckout(argInfo.checkoutFile, argInfo.pricesFile)

	if err != nil {
		panic(err)
	}

	fmt.Printf("checkout file: %s, prices file: %s, total value of checkout: %v\n", argInfo.checkoutFile, argInfo.pricesFile, result)
}
