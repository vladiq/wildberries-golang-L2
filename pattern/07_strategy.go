package main

import "fmt"

type Cooker interface {
	Apply() string
}

type Cook struct {
	recipe Cooker
}

func (p *Cook) Apply() string {
	return p.recipe.Apply()
}

type Pepperoni struct {
	doughWeight     int
	numTomatoes     int
	blackPepperG    int
	oregano         int
	pepperoniSlices int
}

func (p *Pepperoni) Apply() string {
	return fmt.Sprintf("Cooking pepperoni with %dg of dough, "+
		"%d tomatoes, %d grams of black pepper, %d oregano leaves and %d pepperoni slices!",
		p.doughWeight, p.numTomatoes, p.blackPepperG, p.oregano, p.pepperoniSlices)
}

type Margarita struct {
	doughWeight int
	numTomatoes int
	mozzarellaG int
	basil       int
}

func (m *Margarita) Apply() string {
	return fmt.Sprintf("Cooking pepperoni with %dg of dough, %d tomatoes, "+
		"%d grams of mozzarella and %d basil leaves!", m.doughWeight, m.numTomatoes, m.mozzarellaG, m.basil)
}

//func main() {
//	pepperoni := &Pepperoni{
//		doughWeight:     150,
//		numTomatoes:     3,
//		blackPepperG:    2,
//		oregano:         3,
//		pepperoniSlices: 10,
//	}
//
//	margarita := &Margarita{
//		doughWeight: 200,
//		numTomatoes: 2,
//		mozzarellaG: 200,
//		basil:       3,
//	}
//
//	margaritaCooker := Cook{margarita}
//	pepperoniCooker := Cook{pepperoni}
//
//	fmt.Println(margaritaCooker.Apply())
//	fmt.Println(pepperoniCooker.Apply())
//}
