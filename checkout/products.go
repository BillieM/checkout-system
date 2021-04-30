package checkout

type (
	CheckoutLine struct {
		Code     string
		Quantity int
	}
	Product struct{}
)

func (cL CheckoutLine) GetCheckoutLinePrice() (int, error) {
	return 0, nil
}

func GetCheckoutPrice(cLSlice []CheckoutLine) (int, error) {

	sum := 0

	for _, cL := range clSlice {

	}

	return sum, nil
}
