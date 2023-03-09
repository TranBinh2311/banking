package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD:
		return true
	case EUR:
		return true
	case CAD:
		return true
	default:
		return false
	}
}
