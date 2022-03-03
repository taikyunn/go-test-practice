package calc

type Sum struct{}

func (a *Sum) SumMulti(n ...int) int {
	var sum int
	for _, m := range n {
		sum += m
	}

	return sum
}
