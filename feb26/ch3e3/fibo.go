package fibonacci

var fibCache = map[uint]uint{0: 0, 1: 1}

func Fib(n uint) uint {
	if val, ok := fibCache[n]; ok {
		return val
	}

	val := Fib(n-1) + Fib(n-2)

	fibCache[n] = val
	return val
}
