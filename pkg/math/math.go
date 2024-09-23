package math

func Gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return Gcd(y, x%y)
}

func LcmOfArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	lcm := arr[0]
	for _, val := range arr {
		x := lcm
		y := val
		gcd := Gcd(x, y)
		lcm = (lcm * val) / gcd
	}
	return lcm
}
