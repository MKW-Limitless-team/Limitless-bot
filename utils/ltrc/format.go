package ltrc

type Format struct {
	Name          string
	ScalingFactor int
	TeamSize      int
	Modifiers     []Modifier
}

type Modifier struct {
	Name     string
	Increase float64
	Decrease float64
}
