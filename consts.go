package inn

const (
	legalEntityINNDigitsCount            = 10
	naturalPersonINNDigitsCount          = 12
	individualEntrepreneurINNDigitsCount = 12
)

var (
	naturalPersonFirstChecksumWeights  = []int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	naturalPersonSecondChecksumWeights = []int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

	individualEntrepreneurFirstChecksumWeights  = []int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	individualEntrepreneurSecondChecksumWeights = []int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

	legalEntityChecksumWeights = []int{2, 4, 10, 3, 5, 9, 4, 6, 8}
)

var (
	naturalPersonForbiddenINN          = "000000000000"
	individualEntrepreneurForbiddenINN = "000000000000"
	legalEntityForbiddenINN            = "0000000000"
)
