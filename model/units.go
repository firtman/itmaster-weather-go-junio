package model

type DegreeK float64
type DegreeC float64
type Meters float64

func (d DegreeK) ToCelsius() DegreeC {
	return DegreeC(float64(d) - 273.15)
}
