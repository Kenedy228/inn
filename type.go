package inn

type Type string

var (
	TypeIP           Type = "ip"
	TypePhysical     Type = "physical"
	TypeOrganization Type = "organization"
)

func (t Type) isValid() bool {
	switch t {
	case TypeIP, TypePhysical, TypeOrganization:
		return true
	default:
		return false
	}
}

func (t Type) codeLength() int {
	switch t {
	case TypeIP, TypePhysical:
		return 12
	case TypeOrganization:
		return 10
	default:
		return 0
	}
}

func (t Type) checksumCoefficients() [][]int {
	switch t {
	case TypeIP, TypePhysical:
		return [][]int{
			{7, 2, 4, 10, 3, 5, 9, 4, 6, 8},
			{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8},
		}
	case TypeOrganization:
		return [][]int{
			{2, 4, 10, 3, 5, 9, 4, 6, 8},
		}
	default:
		return [][]int{}
	}
}
