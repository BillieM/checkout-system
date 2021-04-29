package checkout

type (
	CheckoutLine struct {
		Code     string
		Quantity int
	}
	Product struct{}
)

func (cL CheckoutLine) GetCheckoutLinePrice() int {
	return 0
}

func GetCheckoutPrice(cLSlice []CheckoutLine) int {
	return 0
}
