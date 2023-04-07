package three

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func Three() {
	s := fmt.Sprintf("%v %v %v", x, y, z)

	fmt.Println(s)
}
