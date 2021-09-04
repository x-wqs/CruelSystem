## Go

- `package main` - entry point
- `GOPATH` - workspace. `GOROOT` - location go installed
- `go build main.go` - compile. `./main.go` run the binary
- `go run main.go` - run the program

### Variables
```golang
var a int
var a = 1 // automatically assigned int type
message := "hello world"
var b, c int = 2, 3
```

### data types

- complext type `var x complex128 = cmplx.Sqrt(-5 + 12i)`
- Arrays have a fixed length defined at declaration, cannot be expanded. `var a[5]int` 
- `var multiD [2][3]int` multi-dimensional array
- slices can be expanded. `var b []int` `numbers := make([]int, 5, 10)`
- slices are an abstraction to an arry. A slice contains: capacity, length, pointer `numbers = append(numbers, 1,2,3,4)`
- maps `var m map[string]int`
- type casting
```golang
a := 1.1
b := int(a)
fmt.Println(b) // -> 1
```

### conditional statements

- if else
```golang
if num := 9; num < 0 {
	//
} else if num < 10 {
	//
} else {
	//
}
```

- switch
```golang
i := 2
switch i {
case 1:
	/* code */
case 2:
	//
default:
	/* code */
	return
}
```

- loop
```golang
for i := 0; i < 10; i ++ {

}
```

- infinite loop 

```golang
for {}
```

- pointers

```golang
var ap *int
a := 12
ap = &a
fmt.Println(*ap) // => 12
```

```golang
func increment(i *int) {
	*i++
}

func main() {
	i := 10
	increment(&i)
	fmt.Println(i) // => 11
}
```

- return of a funciton being predefined
```golang
func add(a int, b int) (c int) {
	c = a + b
	return
}

func main() {
	fmt.Println(add(2, 1))
} // => 3
```

- return multiple values
```golang
func add(a int, b int) (int, string) {
  c := a + b
  return c, "successfully added"
}
func main() {
  sum, message := add(2, 1)
  fmt.Println(message)
  fmt.Println(sum)
}
```

### struct

- struct is a typed, collection of different fields, group data together

```golang
type person struct {
  name string
  age int
  gender string
}
```

```golang
pp = &person{name: "Bob", age: 42, gender: "Male"}
pp.name
```
