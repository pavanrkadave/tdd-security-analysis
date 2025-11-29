package valueobjects

type Likelihood int

const (
	Rare Likelihood = iota
	Occasional
	Frequent
)

func (likelihood Likelihood) String() string {
	switch likelihood {
	case Rare:
		return "RARE"
	case Occasional:
		return "OCCASIONAL"
	case Frequent:
		return "FREQUENT"
	default:
		return "UNKNOWN"
	}
}
