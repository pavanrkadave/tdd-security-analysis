package valueobjects

type Measure struct {
	description string
}

func NewMeasure(description string) *Measure {
	return &Measure{description: description}
}

func (measure *Measure) Description() string {
	return measure.description
}
