package even

func IsEven(n int) bool {
	return n%2 == 0
}

type Kilometers float64

func (km Kilometers) InMiles() float64 {
	return float64(km * 0.621371)
}
