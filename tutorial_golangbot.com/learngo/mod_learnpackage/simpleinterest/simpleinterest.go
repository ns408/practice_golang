package simpleinterest

//Calculate calculates and returns the simple interest for a principal p, rate
//of interest r for duration t years
//func Calculate(p float64, r float64, t float64) float64 {
func Calculate(p, r, t float64) (interest float64) { // - using Named return value
	//interest := p * (r / 100) * t // - only works when not using Named return value
	interest = p * (r / 100) * t
	return interest
}
