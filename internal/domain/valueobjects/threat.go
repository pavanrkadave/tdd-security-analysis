package valueobjects

type Threat struct {
	name string
}

func NewThreat(name string) *Threat {
	return &Threat{
		name: name,
	}
}

func (threat *Threat) Name() string {
	return threat.name
}
