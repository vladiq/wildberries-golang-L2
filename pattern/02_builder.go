package main

type pizzaBuilder struct {
	toppings []string
	sauce    string
	dough    string
	diameter int
}

func (pb *pizzaBuilder) AddTopping(topping string) *pizzaBuilder {
	pb.toppings = append(pb.toppings, topping)
	return pb
}

func (pb *pizzaBuilder) SetSauce(sauce string) *pizzaBuilder {
	pb.sauce = sauce
	return pb
}

func (pb *pizzaBuilder) SetDough(dough string) *pizzaBuilder {
	pb.dough = dough
	return pb

}

func (pb *pizzaBuilder) SetDiameter(diameter int) *pizzaBuilder {
	pb.diameter = diameter
	return pb
}

func NewPizzaBuilder() *pizzaBuilder {
	return &pizzaBuilder{}
}

type Director struct {
	pizzaBuilder *pizzaBuilder
}

func (d *Director) buildPepperoni() *pizzaBuilder {
	return NewPizzaBuilder().AddTopping("Mozzarella").AddTopping("Sausage").AddTopping("Chili pepper").AddTopping("Tomatoes").SetDough("Thin").SetDiameter(33)
}

func (d *Director) buildСapricciosa() *pizzaBuilder {
	return NewPizzaBuilder().AddTopping("Mozzarella").AddTopping("Ham").AddTopping("Mushrooms").AddTopping("Tomatoes").SetDough("Deep Dish").SetDiameter(25)
}

func NewDirector() *Director {
	return &Director{}
}

//func main() {
//	director := NewDirector()
//	pizza := director.buildPepperoni()
//	fmt.Println(pizza)
//	pizza = director.buildСapricciosa()
//	fmt.Println(pizza)
//
//	pizza = NewPizzaBuilder().AddTopping("Meat").AddTopping("Cheese").SetDough("Thin crust").SetDiameter(100)
//	fmt.Println(pizza)
//}
