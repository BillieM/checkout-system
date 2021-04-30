/*
Package checkout provides ...

*/
package checkout

type (
	/*
		CheckoutLine stores
	*/
	CheckoutLine struct {
		Code     string
		Quantity int
	}
	Product struct {
		Price int
		Offer map[int]int
	}
)

func (cL CheckoutLine) GetCheckoutLinePrice() (int, error) {
	return 0, nil
}

func GetCheckoutPrice(cLSlice []CheckoutLine) (int, error) {

	sum := 0

	for _, cL := range cLSlice {
		_ = cL
	}

	return sum, nil
}
