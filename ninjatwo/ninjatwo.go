package ninjatwo

import "fmt"

const (
	myTypedConst   int8 = 12
	myUntypedConst      = "hey"
)

func NinjaTwo() {
	first()
	second()

	fmt.Println(myTypedConst, myUntypedConst)

	third()
	four()
	five()
}

func first() {
	myNumber := 42

	fmt.Printf("%d decimal\n%b binary\n%x hexadecimal", myNumber, myNumber, myNumber)
}

func second() {
	conditional := 1 == 1
	conditional1 := 1 <= 5
	conditional2 := 5 >= 1
	conditional3 := true != false
	conditional4 := 4 < 2
	conditional5 := len("Hey") > len("Hi")

	fmt.Println(conditional)
	fmt.Println(conditional1)
	fmt.Println(conditional2)
	fmt.Println(conditional3)
	fmt.Println(conditional4)
	fmt.Println(conditional5)

}

func third() {
	hey := 42

	fmt.Printf("%d decimal\n%b binary\n%x hexadecimal\n", hey, hey, hey)

	heyButShifted := 42 << 1
	fmt.Printf("%d decimal\n%b binary\n%x hexadecimal\n", heyButShifted, heyButShifted, heyButShifted)
}

func four() {
	myString := `hey`

	fmt.Printf("%v\n", myString)
}

const (
	_          = iota
	firstYear  = iota + 2023
	secondYear = iota + 2023
	thirdYear  = iota + 2023
	fourYear   = iota + 2023
)

func five() {
	fmt.Println(firstYear, secondYear, thirdYear, fourYear)
}
