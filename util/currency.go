package util

var supportedCurrencies = []string{
	"USD",
	"EUR",
	"CAD",
}

func IsSupportedCurrency(currency string) bool {
	for _, curr := range supportedCurrencies {
		if curr == currency {
			return true
		}
	}
	return false
}
