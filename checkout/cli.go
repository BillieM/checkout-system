package checkout

import (
	"flag"
	"fmt"
	"io"
	"os"
)

// Constants CheckoutPath and ProductsPath serve as default paths to JSON data files should they not be given.
const (
	// Default checkout data filePath
	CheckoutPath = "./checkout_data.json"

	// Default price data filePath
	ProductsPath = "./product_data.json"
)

//ArgInfo is returned from getArgInfo and contains the filepaths for the checkout file/ products file (if they are given)
//
//Filepaths may be relative or absolute
type ArgInfo struct {
	checkoutPath string // checkout json file path
	productsPath string // products json file path
}

// getArgInfo returns an instance of ArgInfo.
//
// if the checkout info file path has not been given, or the products flag has not been given,
// checkoutFile/ productsFile will be returned as "" respectively.
//
//Filepaths may be relative or absolute.
func GetArgInfo() ArgInfo {

	// get products flag value for products file
	productsPath := flag.String("products", ProductsPath, "optional filepath to products JSON")
	flag.Parse()

	// get first positional argument for checkout file
	checkoutPath := flag.Arg(0)

	// set to default CheckoutPath if flag empty/ checkout argument not given

	if checkoutPath == "" {
		checkoutPath = CheckoutPath
	}

	return ArgInfo{
		checkoutPath: checkoutPath,
		productsPath: *productsPath,
	}
}

// CheckoutCLI is called from the parent main package, and is the primary entry point.
// It accepts an io.writer to write the result string to.
//
// CLI command takes a filename as an argument, expecting a json file of checkout lines,
// If no filename argument is given, the default checkout dataset will instead be used.
//
// An optional products flag can also be given to specify a path to a different products list.
func CheckoutCLI(out io.Writer) error {
	// --help info
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] <inJSONLocation>\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	argInfo := GetArgInfo()

	// logic to extract from json/ calc checkout value
	result, err := ProcessCheckout(argInfo.checkoutPath, argInfo.productsPath)

	if err != nil {
		return err
	}

	fmt.Fprintf(out, "checkout file: %s\nproducts file: %s\ntotal value of checkout: %v\n", argInfo.checkoutPath, argInfo.productsPath, result)

	return nil
}
