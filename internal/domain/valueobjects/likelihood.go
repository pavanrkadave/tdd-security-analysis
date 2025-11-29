package valueobjects

type Likelihood int

const (
	Rare Likelihood = iota
)

func (likelihood Likelihood) String() string {
	return "RARE"
}
