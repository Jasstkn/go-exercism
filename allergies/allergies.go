package allergies

var allergiesMapping = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) (out []string) {
	for k := range allergiesMapping {
		if AllergicTo(allergies, k) {
			out = append(out, k)
		}
	}
	return out
}

func AllergicTo(allergies uint, allergen string) bool {
	return allergies&allergiesMapping[allergen] != 0
}
