package Chapter02

type Pound float64
type Kilo float64

func PoundToKilo(p Pound) Kilo {
	return Kilo(p * 0.45)
}
