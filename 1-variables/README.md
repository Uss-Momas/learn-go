# Variables In GO
In Go, _variables_ are explicitly declared and used by the compiler to e.g. check type-correctness of function calls.

The ```var``` statement declares a list of variables; as in function argument lists, the type is last.

A ```var``` statement can be at package or function level. We see both in this example.

example:
```go
package main

import "fmt"

var c, python, java bool // initial value will be false

func main() {
    var i int // initial value will be 0
    fmt.Println(i, c, python, java) // 0, false, false, false
}
```

The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. <br>This syntax is only available inside functions.

```go
var f = "apple"
b := "banana"
```

## Constants
keyword: ```const variable_name type = "VALUE"```

Example:

```go
package main

import "fmt"
import "math"

const s string = "constant"

func main() {

    fmt.Println(s)

    const n = 5000000000

    const d = 3e20 / n
    fmt.Println(d)
}
```