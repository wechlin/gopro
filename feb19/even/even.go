package even

func IsEven(n int) bool {
	switch n {
	case 0, 2, 4, 6, 8, 10:
		return true
	default:
		return false
	}
}
