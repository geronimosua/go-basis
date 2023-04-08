package ninjathree

import (
	"fmt"
)

func Ninjathree() {

	//first

	for i:= 1; i <= 10000 ; i++ {
		fmt.Println(i)
	}
	
	//second
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\n", i)
		for j := 0; j < 3; j++{
			fmt.Printf("\t%#U\n", i)
		}
	}


	//third
	yearsAlive := 2003
	for yearsAlive <= 2023 {
		fmt.Printf("%d\n", yearsAlive)
		yearsAlive++
	}

	//four
	for {
		yearsAlive--

		fmt.Printf("years alive without condition in for loop %d\n", yearsAlive)
		if yearsAlive == 2003 {
			break
		}
	}

	//five
	for i := 10; i <= 100; i++ {
		fmt.Printf("%d\n", i%4)
	}

	//six & seven
	fmt.Printf("Is your name Scarlett?\n If yes press 'Y' if no press 'N'\n")
	var isScarlett string 
	fmt.Scan(&isScarlett)

	 if isScarlett == "Y" {
		fmt.Println("Hello scarlett")
	 } else if isScarlett == "N" {
		fmt.Println("Goodbye asshole")
	 }

	 //eight
	 myName := "gopher"
	 switch {
	 case myName == "James Bond":
			fmt.Printf("I'm James Bond\n")
	 case myName == "gopher":
		fmt.Printf("I like carrots so hard\n")
	 }

	//nine
	do(myName)
}

func do(i interface{}) {
switch v := i.(type) {
	case int:
			fmt.Printf("Your name has to be a string")
		case string:
			fmt.Printf("%s! That's a good name! haha\n", i)
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}