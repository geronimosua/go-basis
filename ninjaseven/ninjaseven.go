package ninjaseven

import "fmt"


type Person struct{ 
	name string;
	age int8
}

func changeMe(p *Person) {
	p.name = "Jhon"

	fmt.Println("New name", p.name)

	friend := *p

	friend.name = "Nahuel"
	friend.age = 22

	fmt.Println("Friend", friend)
	fmt.Println("myself", p)
}

func NinjaSeven() {
	name := "Geronimo"

	fmt.Println(&name)

	myself := Person{
		name: name,
		age: 20,
	}

	changeMe(&myself)
}