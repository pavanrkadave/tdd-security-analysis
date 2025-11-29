package valueobjects

type Impact int

const (
	Marginal Impact = iota
	Moderate
	Critical
)

func (impact Impact) String() string {
	switch impact {
	case Marginal:
		return "MARGINAL"
	case Moderate:
		return "MODERATE"
	case Critical:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}
