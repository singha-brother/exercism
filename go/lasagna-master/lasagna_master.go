package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodleCount int = 0
	var sauceCount float64 = 0
	for _, layer := range layers {
		if layer == "noodles" {
			noodleCount += 50
		}
		if layer == "sauce" {
			sauceCount += 0.2
		}
	}
	return noodleCount, sauceCount
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(fList []string, myList []string) {
	secret := fList[len(fList)-1]
	myList[len(myList)-1] = secret
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(qty []float64, p2 int) []float64 {
	newScale := make([]float64, len(qty))
	for i := 0; i < len(qty); i++ {
		newScale[i] = qty[i] * float64(p2) / 2
	}
	return newScale
}
