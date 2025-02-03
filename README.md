### 01. How to start go
- Installing Go: Go installation and GOPATH setup.
- Understanding the Workspace: Explanation of go.mod and GOPATH.
- Basic Commands:
```bash
go mod init project_name
go run file_name.go
go run project_name
and go get -u "repo link"
go build
```

### 02. Data Types and common variable rules
- **Variable Declaration**: 
  - Use `var` for explicit declarations:  
    ```go
    var x int = 10
    ```
  - Use short declaration (`:=`) for concise syntax:  
    ```go
    y := 20
    ```
  - Constants are declared with `const`:  
    ```go
    const pi = 3.14
    ```
- **Zero Values**: 
  - Uninitialized variables are assigned a zero value based on their type:
    - `int`: `0`
    - `float64`: `0.0`
    - `string`: `""`
    - `bool`: `false`

- **Type Inference**: 
  - The `:=` operator infers the type from the assigned value:
    ```go
    z := "Hello" // z is inferred as string
    ```
  - Use `var` when you need to explicitly specify the type.

- **Type Aliases**: 
  - Create custom types or aliases using `type`:
    ```go
    type Celsius float64
    var temperature Celsius = 25.5
    ```
  - Type aliases help improve code readability and enforce type safety.
  
### 03. Statement
- If-Else: Multi-condition handling.
```go
package main
import ("fmt")
func main(){
	var x int;
	fmt.Scan(&x);
	if x%2 == 0 {
		fmt.Println("Even")
	} else if x > 1{
		fmt.Println("Odd")
	} else{
		fmt.Println("Neither")
	}
}
```
- Defer
```go
package main
import "fmt"
func main() {
    defer fmt.Println("This will be printed last")
    fmt.Println("This will be printed first")
}
// This will be printed first
// This will be printed last
```
### 04. Loop
- For Loop with range
```go
fruits := []string{"apple", "banana", "orange"}
for i := range fruits{
	fmt.Println(fruits[i])
}
```
- Index, value
```go
var fruit = []string{"Apple", "Banana", "Watermelon"}
for index, value := range fruit{
	fruit := fmt.Sprintf("index is : %d name is : %s", index, value)
	fmt.Println(fruit)
	// fmt.Printf("Index is : %d value is :%s \n", index, value)
	// %s = string, %d = decimal/int,  %v = value, %t = type, %f =float, %+v = struct with field name
}
```
- Normal Loop in one statement
```go
fruits := []string{"apple", "banana", "orange"}
for i := 0; i<len(fruits); i++{
	fmt.Println(fruits[i])
}
```
- Like while Loop
```go
fruits := []string{"apple", "banana", "orange"}
i := 0
for i < len(fruits){
	fmt.Println(fruits[i])
	i++
}
```
- Infinity loop with break
```go
i := 1
for{
	fmt.Println(i)
	if i == 100{
		break
	}
	i += 1
}
```
- loop with goto
```go
	i := 1
	for{
		fmt.Println(i)
		if i == 100{
			goto finish
		}
		i += 1
	}
	finish:
		fmt.Println("Finished")
```

### 05 Function
- Normal Function
```go
func add(x, y int) int {  // func name(arg1, arg2, type) return_type { area}  return type can be multiple (int, string)
	return x + y
}
func main() {
	result := add(1, -1)
	fmt.Println(result)
}
```
- Function with error handle
```go
package main
import (
	"errors"
	"fmt"
)
func add(x, y int) (int, error) {  // return errpr type with value
	if x < 0 || y < 0 {
		return 0, errors.New("x or y is less than 0")
	}else{
		return x + y, nil
	}
}
func main() {
	// result, _ := add(1, -1)   // if receive only value then  use _ instead or err or this kind
	// fmt.Println(result)

	result, err := add(1, 2)
	if err != nil {
		fmt.Println(err)	
	}else{
		fmt.Println(result)
	}
}
```
- Pass unkwon multiple arguments argument
```go
func calculate(a ...int) int {
	sum := 0
	for i := range a {
		sum += a[i]
	}
	return sum
}
func main() {
	lsit := []int{1, 2, 3, 4, 5}
	result := calculate(lsit...)
	// result = calculate(1, 2, 3, 4, 5)
	fmt.Println(result)
}
```
- Pass array as argument
```go
func calculate(a []int) int {
	sum := 0
	for i := range a {
		sum += a[i]
	}
	return sum
}

func main() {
	lsit := []int{1, 2, 3, 4, 5}
	result := calculate(lsit)
	fmt.Println(result)
}
```


### 06. call back function
- Passing another Functions & Anonymous function
```go
package main

func calculate(x, y int, callback func(int)int) int{
	result := x + y
	return callback(result)
}
func add(x int) int{
	return x * x
}

func main() {
	x, y := 1, 2
	result1 := calculate(x, y, add)  // call another function as a parameter
	result2 := calculate(x, y, func(x int) int{  // call anonymous function, rules is same amount and type args invoked callback
		return x * x * x
	})
	println(result1)
	println(result2)
}
```
### 07. Common Struct and method 
- Simple Struct
```go
package main
import "fmt"
type Student struct {
	id   int
	name string
	GPA  float32
}
func main() {
	// s1 := Student{1, "John", 3.93}
	s1 := Student{id : 1, name : "John", GPA: 3.93}  // best way dosen't give error less args input
	fmt.Println(s1.id)
	s1.id = 2 // change the value of id
	fmt.Println(s1.name)
	fmt.Println(s1.GPA)
	fmt.Println(s1.id)

}
```
- Methods
```go
package main
import "fmt"
type Student struct {
	id   int
	name string
	number  int
	GPA float32
}

func (s *Student) calculateGPA(){   // if dosen't pass as pointer then main struct never update it will return just a copy
	if s.number > 80 {
		s.GPA = 4.0
	}else if s.number > 70 {
		s.GPA = 3.0
	}else {
		s.GPA = 2.0
	}
}

func main() {
	s1 := Student{id : 1, name :"John", number: 75}
	s1.calculateGPA()
	fmt.Println(s1.GPA)
}
```
- Handle slice/array with struct 
```go
func main() {
	// Create a slice of students
	students := []Student{
		{id: 1, name: "Samrat Biswas", number: 60},
		{id: 2, name: "John Doe", number: 85},
		{id: 3, name: "Jane Smith", number: 75},
		{id: 4, name: "Alice Johnson", number: 65},
	}

	// Calculate GPA for each student
	for i := range students {
		students[i].calculateGpa()
	}

	// Print the details of all students
	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s, Number: %d, GPA: %.1f\n",
			student.id, student.name, student.number, student.GPA)
	}
}
```
### 08. Common Built-in methods
- fmt
```go
fmt.Print("Hello, ")
fmt.Print("World!")
// Output: Hello, World!

fmt.Println("Hello, World!")
fmt.Println("Welcome to Go!")
// Hello, World!
// Welcome to Go!

name := "Alice"
age := 30
fmt.Printf("Name: %s, Age: %d\n", name, age)
// Output: Name: Alice, Age: 30

name := "Bob"
age := 25
s := fmt.Sprintf("Name: %s, Age: %d", name, age)
fmt.Println(s)
// Output: Name: Bob, Age: 25

s := fmt.Sprintln("Hello, World!")
fmt.Print(s)
// Output: Hello, World!

err := fmt.Errorf("invalid value: %d", 42)
fmt.Println(err)
// Output: invalid value: 42

var name string
var age int
fmt.Print("Enter your name and age: ")
fmt.Scan(&name, &age)
fmt.Printf("Name: %s, Age: %d\n", name, age)
// Input: Alice 30
// Output: Name: Alice, Age: 30

var name string
var age int
fmt.Print("Enter your name and age: ")
fmt.Scanf("%s %d", &name, &age)
fmt.Printf("Name: %s, Age: %d\n", name, age)
// Input: Alice 30
// Output: Name: Alice, Age: 30

var name string
var age int
fmt.Print("Enter your name and age: ")
fmt.Scanln(&name, &age)
fmt.Printf("Name: %s, Age: %d\n", name, age)
// Input: Alice 30
// Output: Name: Alice, Age: 30

// fmt.Printf("Index is : %d value is :%s \n", index, value)
// %s = string, %d = decimal/int,  %v = value, %t = type, %f =float, %+v = struct with field name, %t = boolean, %p = pointer, %b = binrar
```
- reflect
```go

```
- Time (Now, Sleep, After, Tick)
```go

```
- len()
```go

```
- cap()
```go

```
- make()
```go

```
- new()
```go

```
- copy()
```go

```
- append()
```go

```
- delete()
```go

```
- close()
```go

```
- panic()
```go

```
- recover()
```go

```
- math
```go

```
- strings
```go

```
- sort
```go

```


### 09. String manipulation
- String Operations:
- - Concatenation and splitting.
- - Trimming, prefix, and suffix checks.
- String Formatting (fmt.Sprintf)
- Regular Expressions
- String to other types (strconv)

### 10. List, Slice, Map, Struct also project base usecase 
- List - with all manipulation's methods (loop, append, delete, pop etc )
- Slice - with all manipulation's methods (loop, append, delete, pop etc )
- Map - with all manipulation (loop, insert key value, remove etc)
```go
func main() {
	students := make(map[string] string)
	students["name"] = "John"
	students["age"] = "20"
	students["GPA"] = "3.14"
	students["isStudent"] = "true"

	for key, value := range students {
		println(key, value)
	}
}
```
- Struct - with all manipulation (nested structs, loop, append, remove etc)
Project Use Cases (which collection/sequence data type where should use)

### 11. OOP Concept with Interface and struct
- Structs and Interfaces:
- - Using structs as objects.
- - Interfaces for polymorphism.
- Composition over Inheritance:
-- Embedding structs.
- Method Receivers:
- - Pointer vs. value receivers.
- Design Patterns:
- - Factory, Singleton, and Adapter in Go.

### 12. Threading/goroutine, Basic Request, Data Scrapping etc
- Threading:
- - Using goroutines for lightweight concurrency.
- - Synchronization with sync.WaitGroup and sync.Mutex.
- Requests:
- - Basic HTTP requests with net/http.
- - Parsing responses (io.Reader).
- Data Scraping:
- - Using colly for web scraping.
- - HTML parsing with goquery.

### 13. Error Handling (Try Except)
- Simple error handeling
```go
package main
import (
	"fmt"
	"strconv"
)
func convert_number(n string) (int64, error){        // int8, 16, 32, 64
	result, err := strconv.ParseInt(n, 10, 64)   // 8, 16, 32, 64
    if err != nil{
		return 0, err
	}else{
		return result, nil      // return int8(result), int16(result), int32(result),  int32(result)/result, 
	}
}

func main(){
	// x,_ := convert_number("10a")
	// fmt.Println(x)
	x, err := convert_number("10a")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(x)
	}
}
```
- Custom Dynamic Error Handle - to integer convert
```go
package main

import (
	"fmt"
	"strconv"
)
func to_integer(n interface{}) (int32, error){

	switch v := n.(type){
	case int:
		return int32(v), nil
	
	case float32:
		return int32(v), nil
	case string:
		result, err := strconv.ParseInt(v, 10, 32)
		if err != nil{
			return 0, err
		}else{
			return int32(result), nil
		}
	default:
		return 0, fmt.Errorf("Unsupported assigned")

	}
}

func main(){
	// x,_ := to_integer("10a")
	// fmt.Println(x)
	x, err := to_integer("10a")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(x)
	}
}

```
- Dynamic error handle - to string convert
```go
package main
import (
	"fmt"
	"strconv"
)
func to_string(number interface{}) (string, error) {
	switch value := number.(type) {
	case int8, int16, int32, int64, int:
		// return fmt.Sprintf("%d", value), nil
		return strconv.FormatInt(int64(value.(int)), 10), nil
	case float32, float64:
		return strconv.FormatFloat(value.(float64), 'f', 2, 64), nil
	case string:
		return value, nil
	default:
		return "", fmt.Errorf("unsupported type: %t", value)
	}
}
func main() {
	x, err := to_string(42)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(x)
	}
}
```
- Dynamic error handle - to Float convert
```go
package main
import (
	"fmt"
	"strconv"
)
func to_float(number interface{}) (float64, error) {
	switch value := number.(type) {
	case int8, int16, int32, int64, int:
		return float64(value.(int)), nil
	case string:
		result, err := strconv.ParseFloat(value, 64)
		if err != nil{
			return 0, err
		}else{
			return result, nil
		}
	case float32, float64:
		return value.(float64), nil
	default:
		return 0, fmt.Errorf("unsupported value %t", value)
	}
}

func main() {

	x, err := to_float(102)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("%.1f", x)
	}

}
```

### 13. File Handling and IO
- File Operations:
- - Reading and writing files (os and io/ioutil).
- Buffering:
- - Using bufio for efficient IO.
- File Manipulation:
- - Checking existence, renaming, and deleting files.
  
## 14. Others
- Time
```go
package main
import ("time"
	"fmt")
func main(){

	fmt.Println(time.Now().Local().Format("01-02-2006 Monday 15:04:05"))
	// Printf is a function that can format string and return value and error status
	currentTime, err := fmt.Printf(time.Now().Format("01-02-2006 Monday 15:04:05"))
	if err != nil{
		fmt.Println(currentTime, err)
	}
	// 
	formatTime := fmt.Sprintf(time.Now().Format("01-02-2006 Monday 15:04:05"))
	fmt.Println(formatTime)
}
```
