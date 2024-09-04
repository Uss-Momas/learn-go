# Lesson 1: Hello World in Go

The idea is simple: print "Hello World" to the screen.

**So  how to do this in GoLang?**<br>
simple: we need the fmt (format) package, which is used to format input, output data.

Methods: fmt.Println, fmt.Print, fmt.Printf

file: ```hello.go```
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
Run the file:
```shell
go run hello.go
```

## Import packages

code:
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

we can also use the form:
```go
import "fmt"
import "math"
```

But its recommended to use factored import statement.