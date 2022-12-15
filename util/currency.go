package util

// Costants for all supposrted currecy 
const(
	USD = "USD"
	EUR = "EUR"
	INR = "INR"
)

// IsSupportedCurrecy is true is the currecy is in the list or false otherwise 
func IsSupportedCurrecy(currency string)bool{
	switch currency{
	case USD,EUR,INR:
		return true
	}
	return false
}