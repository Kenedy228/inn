package inn

const (
	legalEntityINNDigitsCount = 10
	personINNDigitsCount      = 12
)

var (
	personFirstChecksumWeights  = []int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	personSecondChecksumWeights = []int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	legalEntityChecksumWeights  = []int{2, 4, 10, 3, 5, 9, 4, 6, 8}
)
