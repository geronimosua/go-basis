package ninjafour

import (
	"fmt"
)

func NinjaFour() {
	
	//1
	first := [5]int{12, 13, 14, 15, 16}

	fmt.Println(first)

	//2
	second := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for _, v := range second {
		fmt.Println(v)
	}
	
	fmt.Printf("Type %T\n", second)

	//3
	fmt.Println(second[:5])
	fmt.Println(second[5:])
	fmt.Println(second[2:7])
	fmt.Println(second[1:6])

	//4

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	otherSlice := []int{56, 57, 58, 59 ,60}
	x = append(x, otherSlice...)
	fmt.Println(x)

	//5
	sliceForExcerciseFive := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	sliceForExcerciseFive = append(sliceForExcerciseFive[:3], sliceForExcerciseFive[6:]...)
	fmt.Println(sliceForExcerciseFive)

	//6
	states := make([]string, 0, 50)

	states = append(states, "Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee"," Texas"," Utah"," Vermont", "Virginia"," Washington", "West Virginia", "Wisconsin", "Wyoming" )

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}

	//7
	james := []string{"James", "Bond", "Shaken, not stirred"}
	missMoneyPenny := []string{"Miss", "Moneypenny", "Helloooo, James"}
	names := [][]string{james, missMoneyPenny}

	for _, v := range names {
		for _, sv := range v {
			fmt.Println(sv)
		}
	}

	//8

	personJames := []string{"Shaken, not stirred", "Martinis", "Women"}
	personMiss := []string{"James Bond", "Literature", "Computer Science"}
	persons := map[string][]string{"bond_james": personJames, "moneypenny_miss": personMiss, "no_dr": {"Being evil", "Ice cream", "Sunsets"} }

	// 9
	persons["Belgrano"] = []string{"Argentina", "Independence", "Flags"}

	for k, v := range persons {
		fmt.Println("this is what", k, "likes")
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
	
	// 10
	delete(persons, "no_dr")
	for k, v := range persons {
		fmt.Println("this is what", k, "likes")
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
}