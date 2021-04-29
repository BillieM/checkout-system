package checkout

import (
	"encoding/json"
	"io/ioutil"
)

func DecodeCheckoutData(filePath string) ([]CheckoutLine, error) {

	/*
		DecodeCheckoutData takes a filepath,
		returns a slice of instances of the CheckoutLine struct and any errors
	*/

	byteArr, err := ioutil.ReadFile(filePath)

	if err != nil {
		return []CheckoutLine{}, err
	}

	cLSlice := []CheckoutLine{}

	err = json.Unmarshal(byteArr, &cLSlice)

	if err != nil {
		return []CheckoutLine{}, err
	}

	return cLSlice, nil
}

func DecodePriceData(filePath string) map[string]Product {

	/*
		DecodePriceData takes a filePath,
		returns a map of product codes to instances of the Product struct and any errors
	*/

	return make(map[string]Product)
}
