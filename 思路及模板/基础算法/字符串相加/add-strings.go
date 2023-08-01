package addStrings

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	res := []byte{}

	for c := 0; i >= 0 || j >= 0 || c >= 0; i, j = i-1, j-1 {
		if i >= 0 {
			c += int(num1[i] - '0')
		}
		if j >= 0 {
			c += int(num2[j] - '0')
		}

		res = append(res, byte(c%10+'0'))
		c /= 10
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}
