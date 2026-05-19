package inn

const (
	legalEntityINNDigitsCount            = 10
	naturalPersonINNDigitsCount          = 12
	individualEntrepreneurINNDigitsCount = 12
)

var (
	personFirstChecksumWeights  = []int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	personSecondChecksumWeights = []int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

	individualEntrepreneurFirstChecksumWeights  = []int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	individualEntrepreneurSecondChecksumWeights = []int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

	legalEntityChecksumWeights = []int{2, 4, 10, 3, 5, 9, 4, 6, 8}
)
