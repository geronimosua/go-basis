package ninjasix

import (
	"fmt"
	"math"
)

type Person struct {
	first string;
	last string;
	age int8;
}

type Circle struct{
	radius float64
}

type Square struct{
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	Area() float64
}

func info(sh Shape) {
	fmt.Println(sh.Area())
}

func NinjaSix() {
	fmt.Println(foo())
	fmt.Println(bar())

	x := []int{2, 3, 4, 5, 6, 7}

	fmt.Println(foo2(x...))
	fmt.Println(bar2(x))

	myFirstDefer()

	me := Person{
		first: "Geronimo",
		last: "Suarez",
		age: 20,
	}
	me.speak()

	myCircle := Circle{
		radius: 2.4,
	}
	mySquare := Square{
		side: 2,
	}

	info(myCircle)
	info(mySquare)

	func(){fmt.Println("I am an anonymus func muahaha")}()

	funcInAVar := foo
	fmt.Println(funcInAVar())

	myVarWithAFuncWhichReturnsAFunc := funcWhichReturnsAFunc()
	fmt.Println(myVarWithAFuncWhichReturnsAFunc()) 

	fmt.Println(squareFirstFoo(foo))

	closureValue := myFirstGoClosure()

	fmt.Println(closureValue())
	fmt.Println(closureValue())


}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 13, "hello"
}

func foo2(x ...int) (sum int) {
	for _, v := range x {
		sum += v
	}
	return
}

func bar2(x []int) (sum int) {
	for _, v := range x {
		sum += v
	}
	return
}

func myFirstDefer() {
	defer fmt.Println("IM DEFERRED :(")
	fmt.Println("That means that I ran first hehe")
}

func (p Person) speak() {
	fmt.Printf("Hi! my name is %v %v and I'm %v years old\n", p.first, p.last, p.age)
}

func funcWhichReturnsAFunc() func() int {
	return func() int {
		return 23
	}
}

func squareFirstFoo(firstFoo func() int) int {
	return firstFoo() * firstFoo()
}

func myFirstGoClosure() func() int {
	var x int

	return func () int {
		for i := 0; i < 13; i++ {
			x += i
		}
		return x
	}
}