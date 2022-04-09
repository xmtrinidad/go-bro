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