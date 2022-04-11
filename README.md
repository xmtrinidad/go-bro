# Go Bro

Using this repo to track my progress as I mess with Go

## Starting a Go project

First need to initialize a project
```
go mod init project-name
```

This will create a go.mod file

### Simple Hello World Example

```go
package main

import "fmt"

func main() {
	fmt.Print("Hello World")
}
```

Then run the Go file with:

```
go run main.go
```

## Functions/Methods I've used so far

```go
fmt.Println("Log")
fmt.Printf("memeMan is %T and totalMemes is %T\n", memeName, totalMemes)
fmt.Printf("Wecome %v you have made %v total memes\n", memeName, totalMemes)

// Length of array/slice
len(arr)
// Append to an array/slice 
append(arr, "Meme 1", "Meme 2", "Meme 3")
// Split a string on whitespace
strings.Fields(meme)
// Check if string contains character
strings.Contains(str, "@")