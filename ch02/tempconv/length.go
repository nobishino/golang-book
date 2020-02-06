package tempconv

import "fmt"

type Meter float64
type Feet float64

const (
	FeetPerMeter Feet  = 3.28084
	MeterPerFeet Meter = 0.3048
)

func (meter Meter) String() string {
	return fmt.Sprintf("%g m", meter)
}

func (feet Feet) String() string {
	return fmt.Sprintf("%g feet", feet)
}

func MeterToFeet(meter Meter) Feet {
	return FeetPerMeter * Feet(meter)
}

func FeetToMeter(feet Feet) Meter {
	return MeterPerFeet * Meter(feet)
}
