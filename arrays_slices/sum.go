package arrays_slices

func Sum(numbers []int) (sum int) {
	for _, n := range numbers {
		sum += n
	}
	return
}

func SumAll(slices ...[]int) (sums []int) {
	for _, slice := range slices {
		sums = append(sums, Sum(slice))
	}
	return
}

func SumAllTails(slices ...[]int) (sums []int) {
	for _, slice := range slices {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(slice[1:]))
		}
	}
	return
}
