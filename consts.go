package inn

const (
	legalEntityDigitsCount            = 10
	naturalPersonDigitsCount          = 12
	individualEntrepreneurDigitsCount = 12
)

var (
	naturalPersonFirstWeights  = []int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	naturalPersonSecondWeights = []int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

	individualEntrepreneurFirstWeights  = []int{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	individualEntrepreneurSecondWeights = []int{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}

	legalEntityWeights = []int{2, 4, 10, 3, 5, 9, 4, 6, 8}
)

var (
	naturalPersonForbiddenValue          = "000000000000"
	individualEntrepreneurForbiddenValue = "000000000000"
	legalEntityForbiddenValue            = "0000000000"
)
