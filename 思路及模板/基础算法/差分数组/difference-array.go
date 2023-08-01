package differenceArray

func getModifiedArray(length int, updates [][]int) []int {
	d := make([]int, length+1)

	for _, u := range updates {
		l, r, c := u[0], u[1], u[2]
		d[l] += c
		if r+1 <= length {
			d[r+1] -= c
		}
	}

	for i := 1; i <= length; i++ {
		d[i] += d[i-1]
	}
	return d
}
