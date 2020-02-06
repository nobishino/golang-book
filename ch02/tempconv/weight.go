package tempconv

import "fmt"

type KillGram float64
type Pound float64

const (
	KillGramPerPound KillGram = 0.453592
	PoundPerKillGram Pound    = 2.20462
)

func (kg KillGram) String() string {
	return fmt.Sprintf("%g kg", kg)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g pound", p)
}

func PoundToKillGram(p Pound) KillGram {
	return KillGram(p) * KillGramPerPound
}

func KillGramToPound(k KillGram) Pound {
	return Pound(k) * PoundPerKillGram
}
