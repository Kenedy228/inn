package inn

type Type string

var (
	TypeIndividualEntrepreneur Type = "individual entrepreneur"
	TypeNaturalPerson          Type = "natural person"
	TypeLegalEntity            Type = "legal entity"
)

func (t Type) isValid() bool {
	switch t {
	case TypeIndividualEntrepreneur, TypeNaturalPerson, TypeLegalEntity:
		return true
	default:
		return false
	}
}

func (t Type) codeLength() int {
	switch t {
	case TypeIndividualEntrepreneur, TypeNaturalPerson:
		return 12
	case TypeLegalEntity:
		return 10
	default:
		return 0
	}
}

func (t Type) checksumCoefficients() [][]int {
	switch t {
	case TypeIndividualEntrepreneur, TypeNaturalPerson:
		return [][]int{
			{7, 2, 4, 10, 3, 5, 9, 4, 6, 8},
			{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8},
		}
	case TypeLegalEntity:
		return [][]int{
			{2, 4, 10, 3, 5, 9, 4, 6, 8},
		}
	default:
		return [][]int{}
	}
}
