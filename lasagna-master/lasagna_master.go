package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgLayerTime int) (preparationTime int) {
	if avgLayerTime == 0 {
		avgLayerTime = 2
	}
	preparationTime = len(layers) * avgLayerTime
	return
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, val := range layers {
		if val == "noodles" {
			noodles += 50
		}
		if val == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsRecipe, myRecipe []string) {
	myRecipe[len(myRecipe)-1] = friendsRecipe[len(friendsRecipe)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amountNeeded []float64, numberOfPortions int) (updatedAmount []float64) {
	updatedAmount = make([]float64, len(amountNeeded))
	copy(updatedAmount, amountNeeded)
	for index, val := range updatedAmount {
		updatedAmount[index] = val * 0.5 * float64(numberOfPortions)
	}
	return
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
