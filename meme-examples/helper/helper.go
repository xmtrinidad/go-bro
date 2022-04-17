package helper

import (
	"fmt"
	"strconv"
)

// Capitalizing the first letter of this function makes it an export
func HelperFunction() {
	fmt.Println("This is a helper function in another file")
}

func MapExample() {
	var myMap = make(map[string]string)
	myMap["memeName"] = "5 gum"
	myMap["userName"] = "MEME MAN"
	myMap["misc"] = "hmm"
	myMap["numberOfMemes"] = strconv.FormatInt(int64(10), 10)

	fmt.Printf("Map example: %v\n", myMap)
}
