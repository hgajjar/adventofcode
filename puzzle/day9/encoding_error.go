package day9

type preambleRange struct {
	min int
	max int
}

func NewPreambleRange(max int) preambleRange {
	return preambleRange{0, max}
}

func FindEncodingError(xmasSeries []int, pr preambleRange) int {
	if pr.max == len(xmasSeries) {
		panic("Range out of bound")
	}

	for i := pr.min; i < pr.max; i++ {
		for j := i + 1; j < pr.max; j++ {
			sum := xmasSeries[i] + xmasSeries[j]
			if sum == xmasSeries[pr.max] {
				pr.min++
				pr.max++

				return FindEncodingError(xmasSeries, pr)
			}
		}
	}

	return xmasSeries[pr.max]
}

func FindSumOfEncryptionWeakness(xmasSeries []int, pr preambleRange) int {
	encryptionErr := FindEncodingError(xmasSeries, pr)

	weaknessSeries := findEncryptionWeakness(xmasSeries, pr, encryptionErr)

	return sumMinMaxOfWeaknessSeries(weaknessSeries)
}

func findEncryptionWeakness(xmasSeries []int, pr preambleRange, encryptionErr int) []int {
	if pr.max == len(xmasSeries) {
		panic("Range out of bound")
	}

	for i := pr.min; i < pr.max; i++ {
		for j := i + 1; j < pr.max; j++ {
			if sum(xmasSeries, i, j) == encryptionErr {
				return xmasSeries[i : j+1]
			}
		}
	}

	pr.min++
	pr.max++

	return findEncryptionWeakness(xmasSeries, pr, encryptionErr)
}

func sum(xmasSeries []int, i int, j int) int {
	sum := 0
	for k := i; k <= j; k++ {
		sum += xmasSeries[k]
	}

	return sum
}

func sumMinMaxOfWeaknessSeries(weaknessSeries []int) int {
	find := func(ws []int, pred func(i, j int) bool) (m int) {
		for i, e := range ws {
			if i == 0 || pred(e, m) {
				m = e
			}
		}
		return
	}

	return find(weaknessSeries, func(i, j int) bool { return i < j }) + find(weaknessSeries, func(i, j int) bool { return i > j })
}
