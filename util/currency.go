package util

import _"github.com/go-playground/locales/currency"

const (
	MAD = "MAD"
	USD = "USD" 
	EUR = "EUR"
	CAD = "CAD"
)


func IsSupportedCurrency(currency string) bool {
	switch currency {
	case MAD,USD,EUR,CAD : return true
	}
	return false
}