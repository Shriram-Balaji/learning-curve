package main

import (
	"errors"
	"fmt"
)

// FUNCTIONS
func add(a int, b int) int {
	return a + b
}

func what(a string) (string, string) {
	return a + " what!", "ignored"
}

// variadic  - functions with variable number of argymenets
func variadic(a ...string) string {
	for _, val := range a {
		fmt.Print(val)
	}
	return ""
}

// a named function, that returns another anonymous function, returning an integer
func increment(initialValue int) func() int {
	i := initialValue
	return func() int {
		i++
		return i
	}
}

// structs ( like classes )
type student struct {
	name       string
	rollNumber int
	subjects   []string
}

// structs with receiver types
type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

// INTERFACES in go
// they provide a way to introduce polymorphism in go lang.

type animal interface {
	run() string
	eat() string
}

type dog struct {
	name  string
	speed string
	food  string
}

func (d dog) eat() string {
	return d.name + " eats " + d.food
}

func (d dog) run() string {
	return d.name + " runs at " + d.speed + "km/hr."
}

type cat struct {
	name  string
	speed string
	food  string
}

func (c *cat) run() string {
	return c.name + " runs at " + c.speed + "km/hr."
}

func (c *cat) newSpeed(updatedSpeed string) string {
	c.speed = updatedSpeed
	return c.name + " says: Nope! I run at " + updatedSpeed + "km/hr"
}

func (c *cat) newFood(food string) string {
	oldFood := c.food
	c.food = food
	return c.name + " says: Ew! I dont like " + oldFood + " I need " + c.food
}

type lion struct {
	name  string
	speed string
	food  string
}

func (l *lion) eat() (string, error) {
	return "", errors.New("Nope! The king in the jungle doesnt take orders from you, human")
}

func main() {
	// SLICES
	/*
		* slices work like arrays in Go, but are a bit diffderent. They are key Datatypes in Go
		* slices can be created using the pre-defined make method, which takes in two arguements
			[]type ( []string, []int)
			len ( length of the slice )
			cap ( optional - capacity of the slice)
		* example: s := make([]string, 4)
	*/

	s := make([]string, 3)
	s[0] = "1"
	s[1] = "2"
	fmt.Println("set:", s)
	fmt.Println("get:", s[1])

	c := make([]string, 2)
	copy(c, s) // inbuilt copy method can be used to copy a ds into another

	fmt.Println(c)

	// MAPS
	// syntax: make(map[key-ttype]vaue-type)
	m := make(map[string]int)
	m["2"] = 1

	// delete - deletes kv pair
	delete(m, "2")
	fmt.Println(m["2"])

	// if a key is deleted, even though its deleted, the value returned from the map's key lookup
	// is converted into a zero-valued response ( 0 or "" etc.,)
	// To avoid that, the optional second value returned during lookup specifies, if the value actually exists or not.
	_, pairs := m["2"]
	// using _ as a blank identifier, does not throw a compile time error of throwing values that are not used
	fmt.Println(pairs) // false

	// RANGE
	nums := []int{1, 2, 3, 4} // a slice ( coz the length is not predefined )
	fmt.Println(nums)
	fmt.Println(" -- SLICE RANGE --")
	for index, value := range nums {
		fmt.Println(index, value)
	}

	Map := map[string]string{"a": "b", "c": "d"}
	fmt.Println(" -- MAP RANGE --")
	for key, value := range Map {
		fmt.Println(key, value)
	}

	threePlusOne := add(3, 1)
	fmt.Println(threePlusOne)

	noWhat, _ := what("NO")
	fmt.Println(noWhat)
	// second value returned is ignored

	yes, notIgnored := what("YES")
	fmt.Println(yes, notIgnored)

	// VARIADIC FUNCTIONS - variable args for a function ( like spread )
	hello := variadic("hello", "hi", "three")
	fmt.Println(hello)

	// closures
	fmt.Println(" -- CLOSURES -- ")
	nextInt := increment(0)
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	// STRUCTS
	fmt.Println("-- STRUCTS --")
	jon := student{name: "jon", rollNumber: 123, subjects: []string{"maths", "physics"}}
	fmt.Println(jon.name + " studies")
	for _, subject := range jon.subjects {
		fmt.Println(subject)
	}

	fmt.Println("-- AREA --")
	Rectangle := rect{width: 10, height: 20}
	rectArea := Rectangle.area()
	fmt.Println(rectArea)

	// use pointer ref to avoid mutating struct
	rp := &Rectangle
	fmt.Println(rp.area())

	fmt.Println("-- INTERFACES in GO --")
	// the dog bichon implements the defined interface animal
	// note that there is no explicit "implements" keyword, but instead
	// the go-compiler identifies if methods defined in a struct, with the same values in the interface, implement it.
	// the dog struct methods are defined as value receivers, that can copy over the object,
	// but not modify the object. Dogs are loyal!
	bichon := dog{name: "bichone frise", speed: "10", food: "pedigree"}
	fmt.Println(bichon.run())
	fmt.Println(bichon.eat())

	// cat receivers are usually naughtier, and hence modify anything given to them
	// they are pointer receivers
	kitty := cat{name: "Kitty", speed: "20", food: "milk"}
	fmt.Println(kitty.run())
	fmt.Println(kitty.newSpeed("30"))
	fmt.Println(kitty.newFood("Fish"))
	fmt.Println(kitty.name + " now eats " + kitty.food)

	// Error handling
	simba := lion{name: "Simba", speed: "50", food: "Deer"}
	_, err := simba.eat()
	fmt.Println("ERROR:", err)

}
