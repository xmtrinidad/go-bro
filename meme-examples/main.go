package main

import (
	"fmt"
	// Importing another file uses the module-name/file.go convention
	"meme-example/helper"
	"strings"
)

type MemeData struct {
	username      string
	numberOfMemes int
	isMeme        bool
}

var memeData = MemeData{
	username:      "meme_man",
	numberOfMemes: 1,
	isMeme:        true,
}

// Package level variables
var packageLevelVariable = "Test Thing"

func main() {
	// Shorthand way to declare var and infer type
	memeName := "Meme Man"
	// Short hand doesn't work on const
	const totalMemes = 50
	// uint type as opposed to int type will prevent negative numbers
	var remainingMemes uint = 10

	testFunction()
	fmt.Println()
	// Function from another package
	helper.HelperFunction()
	fmt.Println()
	helper.MapExample()

	introMeme(memeName, totalMemes, remainingMemes)

	var memes []string
	memes = append(memes, "Meme 1", "Meme 2", "Meme 3")
	fmt.Printf("The memes slice looks like this %v\n", memes)
	fmt.Printf("The memes slice length is: %v\n", len(memes))

	for index, meme := range memes {
		fmt.Printf("The index is %v and the meme is %v\n", index, meme)
		// Slicing a string
		var split = strings.Fields(meme)
		fmt.Println(split)
	}

	// for  {
	// 	var userMeme string
	// 	// passing in a pointer
	// 	fmt.Scan(&userMeme)

	// 	// Array stuff
	// 	memes = append(memes, memeName+" created the "+userMeme+" meme")

	// 	remainingMemes = remainingMemes - 1

	// 	fmt.Printf(memes[0])
	// 	fmt.Printf("There are %v memes remaining\n", remainingMemes)

	// 	fmt.Printf("The memes slice looks like this: %v\n", memes)
	// 	fmt.Printf("The memes slice type is: %T\n", memes)
	// 	fmt.Printf("The memes slice length is: %v\n", len(memes))
	// }

	if len(memes) < 4 {
		fmt.Println("Memes slice is less than 3")
	} else {
		fmt.Println("More than 3 memes")
	}

	thing := "Bruh"

	switch thing {
	case "Bro":
		fmt.Printf("Thing is Bro\n")
		break
	case "Bruh":
		fmt.Printf("Thing is Bruh\n")
		break
	default:
		fmt.Println("NO")
	}
}

func testFunction() {
	fmt.Printf("This is a package level variable %v", packageLevelVariable)
	fmt.Println("This is an example function")
}

func introMeme(name string, total int, remaining uint) {
	// Basic Println statements
	fmt.Println("Welcome", name, "you have made", total, "total memes.")
	fmt.Println("There are still", remaining, "memes to make.")
	fmt.Println()

	// Another way to write the above println statements
	fmt.Printf("Wecome %v you have made %v total memes\n", name, total)
	fmt.Printf("There are still %v memes to make.\n", remaining)
}

// example function with type array returning a string
func thing(myArray []string) string {
	return "memes array"
}

// Return multiple values

func returnMultipleThings() (string, int, []string) {
	var things []string
	return "thing", 1, append(things, "test 1", "test2")
}
