package main

import (
	"fmt"
	"strings"
)

func main() {
	// Shorthand way to declare var and infer type
	memeName := "Meme Man"
	// Short hand doesn't work on const
	const totalMemes = 50
	// uint type as opposed to int type will prevent negative numbers
	var remainingMemes uint = 10

	// Example of showing variable types
	fmt.Printf("memeMan is %T and totalMemes is %T\n", memeName, totalMemes)
	fmt.Println()

	// Basic Println statements
	fmt.Println("Welcome", memeName, "you have made", totalMemes, "total memes.")
	fmt.Println("There are still", remainingMemes, "memes to make.")
	fmt.Println()

	// Another way to write the above println statements
	fmt.Printf("Wecome %v you have made %v total memes\n", memeName, totalMemes)
	fmt.Printf("There are still %v memes to make.\n", remainingMemes)

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

}
