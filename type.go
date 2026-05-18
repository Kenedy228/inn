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
