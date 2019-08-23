package space

// Planet should contain name of planet
type Planet string

const rotateTime = 31557600

// Age function calculates age using planet orbital period and time in seconds
func Age(seconds float64, planet Planet) float64 {
	planets := map[Planet]float64{
		"Earth":   rotateTime,
		"Mercury": rotateTime * 0.2408467,
		"Venus":   rotateTime * 0.61519726,
		"Mars":    rotateTime * 1.8808158,
		"Jupiter": rotateTime * 11.862615,
		"Saturn":  rotateTime * 29.447498,
		"Uranus":  rotateTime * 84.016846,
		"Neptune": rotateTime * 164.79132}

	return seconds / (planets[planet])

}
