package checkout

import (
	"encoding/json"
	"io/ioutil"
)

// DecodeCheckoutData takes a filePath and returns a slice of instances of CheckoutLine.
//
// An error is returned if the file cannot be read due to a non-existent file or invalid filePath,
// or if the the files content is not JSON data capable of being being unmarshaled into []CheckoutLine
// (i.e. it must be contain an array of objects with a product code and quantity value)
func DecodeCheckoutData(filePath string) ([]CheckoutLine, error) {

	// read file into byteSlice
	byteSlice, err := ioutil.ReadFile(filePath)

	if err != nil {
		return []CheckoutLine{}, err
	}

	// marshal data from byteSlice into a slice of CheckoutLine
	cLSlice := []CheckoutLine{}
	err = json.Unmarshal(byteSlice, &cLSlice)

	if err != nil {
		return []CheckoutLine{}, err
	}

	return cLSlice, nil
}

// DecodeProductData takes a filePath and returns a map of [productCode]Product.
//
// An error is returned if the file cannot be read but to a non-existent file or invalid filePath,
// or if the files content is not JSON data capable of being unmarshaled into map[string]Product.
// (i.e. it must contain an object using product code strings as keys to another object with Price/ OfferQuantity/ OfferPrice)
func DecodeProductData(filePath string) (map[string]Product, error) {

	// read file into byte slice
	byteSlice, err := ioutil.ReadFile(filePath)

	if err != nil {
		return map[string]Product{}, err
	}

	// marshal data from byteSlice into a map of [prodCodes]Product
	prodMap := map[string]Product{}
	err = json.Unmarshal(byteSlice, &prodMap)

	if err != nil {
		return map[string]Product{}, err
	}

	return prodMap, nil
}
