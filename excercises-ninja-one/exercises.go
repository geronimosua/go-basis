package excercises

import (
	"fmt"

	"github.com/geronimosua/go-basis/excercises-ninja-one/second"
	"github.com/geronimosua/go-basis/excercises-ninja-one/three"
)

type myOwnType int

var x myOwnType

var y int

func Excercises() {

	fmt.Println("First exercise")
	first()
	fmt.Println("##########")

	fmt.Println("Second exercise")
	second.Second()
	fmt.Println("##########")

	fmt.Println("Third exercise")
	three.Three()
	fmt.Println("##########")

	fmt.Println("Four excercise")
	fmt.Printf("X value is %v the type is %T\n", x, x)

	x = 42

	fmt.Printf("And the new value of X is %v\n", x)
	fmt.Println("##########")

	fmt.Println("Fifth excercise")
	fmt.Printf("X value is %v the type is %T\n", x, x)

	x = 13

	fmt.Printf("The value of X is %v\n", x)
	y = int(x)
	fmt.Printf("The value of Y is %v\n", y)
	fmt.Printf("The type of Y is %T\n", y)
	fmt.Println("##########")
}
