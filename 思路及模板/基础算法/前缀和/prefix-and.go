package prefixAnd

type NumArray struct {
	s []int
}

func Constructor(nums []int) NumArray {

	n := len(nums)
	sum := make([]int, n+1)
	for k, v := range nums {
		sum[k+1] = sum[k] + v
	}

	return NumArray{sum}
}

func (n NumArray) SumRange(left, right int) int {
	return n.s[right+1] - n.s[left]
}
