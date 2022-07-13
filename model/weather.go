package model

type Weather struct {
	Temperature DegreeK
	FeelsLike   DegreeK
	Humidity    int
	Condition   int
	Visibility  Meters
}
