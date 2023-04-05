package util

const (
	USD = "USD"
	EUR = "EUR"
	GBP = "GBP"
	BDT = "BDT"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, GBP, BDT:
		return true
	default:
		return false
	}
}
