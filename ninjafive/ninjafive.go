package ninjafive

import "fmt"

func NinjaFive() {

	//1
	type Person struct {
		firstName         string
		lastName          string
		favoriteIceCreams []string
	}

	gero := Person{
		firstName:         "Geronimo",
		lastName:          "Suarez",
		favoriteIceCreams: []string{"Banana split", "Irish cream", "Javanes coffe"},
	}

	agus := Person{
		firstName:         "Agustina",
		lastName:          "Brito",
		favoriteIceCreams: []string{"Switzerland chocolate", "Hailed chocolate"},
	}

	fmt.Println("Gero")
	for _, v := range gero.favoriteIceCreams {
		fmt.Printf("\t%v", v)
	}
	fmt.Printf("\n")
	fmt.Println("Agus")
	for _, v := range agus.favoriteIceCreams {
		fmt.Printf("\t%v", v)
	}
	fmt.Printf("\n")


	//2 
	myMap := map[string]Person{
		gero.lastName: gero,
		agus.lastName: agus,
	}

	for k, v := range myMap {
		fmt.Printf("\n%v \t%v\n", k, v)
	}

	//3

	type vehicle struct {
		doors int8
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	myTruck := truck{
		vehicle: vehicle{doors: 4, color: "white"},
		fourWheel: true,
	}

	mySedan := sedan{
		vehicle: vehicle{doors: 2, color: "blue"},
		luxury: false,
	}

	fmt.Printf("%v\n", myTruck)
	fmt.Printf("%v\n", mySedan)

	fmt.Println("Truck")

	fmt.Println("color",myTruck.color)
	fmt.Println("doors", myTruck.doors)
	fmt.Println("is four wheel", myTruck.fourWheel)
	
	fmt.Println("Sedan")
	fmt.Println("color", mySedan.color)
	fmt.Println("doors", mySedan.doors)
	fmt.Println("is luxury", mySedan.luxury)

	//4

	myGero := struct {
		name string
		topFiveSongs [5]string
		gamesHePlay []string
		notesInHighschool map[string]int
	}{
		name: "Gero",
		topFiveSongs: [5]string{"Cigarettes", "bender++girlfriend", "Shorty", "yesterday", "Baby"},
		gamesHePlay: []string{"League of legends", "Cod"},
		notesInHighschool: map[string]int{
			"Math": 10,
			"History": 10,
			"Geography": 10,
			"English": 8,
		},
	}

	for i, v := range myGero.topFiveSongs {
		fmt.Printf("index %v song %v\n", i, v)
	}

	for i, v := range myGero.gamesHePlay {
		fmt.Printf("index %v game %v\n", i, v)
	}

	for k, v := range myGero.notesInHighschool {
		fmt.Printf("key %v note %v\n", k, v)
	}
}
