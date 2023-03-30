package main

import "fmt"

func main() {
	println("Arrays and Slices in Go")

	// [1, 2, 3, 4]
	planets := [8]string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "nepture"}
	fmt.Println(planets)
	fmt.Println(planets[0:])

	var planetsArray [8]bool
	planetsArray[0] = true
	fmt.Println(planetsArray)

	planetSlice := []string{"mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "nepture"}
	fmt.Println(planetSlice)
	fmt.Println(planetSlice[4])

	var planetSliceVerbose []string
	planetSliceVerbose = append(planetSliceVerbose, "newera")
	fmt.Println(planetSliceVerbose)
}
