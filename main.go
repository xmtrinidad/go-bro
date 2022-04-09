package main

import "fmt"

func main() {
	var memeName = "Meme Man"
	const totalMemes = 50
	var remainingMemes = 10

	fmt.Println("Welcome", memeName, "you have made", totalMemes, "total memes.")
	fmt.Println("There are still", remainingMemes, "memes to make.")
	fmt.Println()

	// Another way to write the above println statements
	fmt.Printf("Wecome %v you have made %v total memes\n", memeName, totalMemes)
	fmt.Printf("There are still %v memes to make.\n", remainingMemes)
}
