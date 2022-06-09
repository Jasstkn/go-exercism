package lasagna

import "fmt"

func PreparationTime(layers []string, averagePreparationTime int) int {
	if averagePreparationTime == 0 {
		averagePreparationTime = 2
	}
	return len(layers) * averagePreparationTime
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodlesPerLayer := 50
	saucePerLayer := 0.2

	for _, i := range layers {
		if i == "noodles" {
			noodles += noodlesPerLayer
		}
		if i == "sauce" {
			sauce += saucePerLayer
		}
	}

	return noodles, sauce
}

func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	fmt.Println(quantities)
	scaledQuantities := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		scaledQuantities[i] = quantities[i] * float64(portions) / 2
	}
	return scaledQuantities
}
