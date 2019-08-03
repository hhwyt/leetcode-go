package main

func reverse(x int) int {
	res := 0

	for x != 0 {
		if math.Abs(float64(res)) > (math.MaxInt32 / 10) {
			return 0
		}
		res = res*10 + x%10
		x /= 10
	}

	return res
}

func main() {

}
