package main

func largestAltitude(gain []int) int {

	altitude := 0
	maxAltitude := 0

	for _, g := range gain {
		altitude += g
		if altitude > maxAltitude {
			maxAltitude = altitude
		}
	}

	return maxAltitude
}

func main() {
	largestAltitude([]int{-4, -3, -2, -1, 4, 3, 2}) //-4 -7 -9 -10 -6 -3 -1
}
