# GO Types

## Basic Types
Go basic types are
- bool
- string
- int int8 int16 int32 int64
- uint uint8 uint16 unit32 uint64 uintptr
- byte // alias for uint8
- rune // alias for int32
- float32 float64
- complex64 complex128


The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.

## Type Convertion

Some numeric conversions:
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
OR
```go
i := 42
f := float64(i)
u := uint(f)
```

Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.