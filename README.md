### 01. How to start go
- Installing Go: Go installation and GOPATH setup.
- Understanding the Workspace: Explanation of go.mod and GOPATH.
- Basic Commands: go run, go build, go test, go fmt, and go get.

### 02. Data Types and common variable rules
### 02. Data Types and Common Variable Rules
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
- Basic for Loop: Iterating over numbers and ranges.
- for with Condition: Using it as a while loop.
- Iterating with range: On slices, maps, and channels.
- Breaking and Continuing: Exiting loops early.

### 05 Function
- Variadic Functions: Functions that accept variable-length arguments.
- Named Return Values: Cleaner return statements.
- Anonymous Functions: Inline functions and closures.
- Method Receiver Functions: Functions tied to structs (key to Go's OOP).

### 06. Common Built-in methods
- Time and Sleep:
- - Formatting and parsing time.
- - Using time.Sleep for delays.
- Type and Typecasting:
- - Explicit casting (int(x)).
- - Checking types with reflection (reflect.TypeOf).
- Slice Operations:
- - append, len, cap, copy.
- Creating Collections:
- - Using make for slices, maps, and channels.

### 07. call back methods
- Passing Functions as Arguments:
- - Defining and using callbacks.
- - Using closures for callback behavior.
- Practical Use Cases:
- - Goroutine callbacks.
- - Sorting using custom comparator callbacks.

### 08. String manipulation
- String Operations:
- - Concatenation and splitting.
- - Trimming, prefix, and suffix checks.
- String Formatting (fmt.Sprintf)
- Regular Expressions
- String to other types (strconv)

### 09. List, Slice, Map, Struct also project base usecase 
- List - with all manipulation's methods (loop, append, delete, pop etc )
- Slice - with all manipulation's methods (loop, append, delete, pop etc )
- Map - with all manipulation (loop, insert key value, remove etc)
- Struct - with all manipulation (nested structs, loop, append, remove etc)
Project Use Cases (which collection/sequence data type where should use)

### 10. Strong OOP in GO (following Java or Python's)
- Structs and Interfaces:
- - Using structs as objects.
- - Interfaces for polymorphism.
- Composition over Inheritance:
-- Embedding structs.
- Method Receivers:
- - Pointer vs. value receivers.
- Design Patterns:
- - Factory, Singleton, and Adapter in Go.

### 11. Threading, Basic Request, Data Scrapping etc
- Threading:
- - Using goroutines for lightweight concurrency.
- - Synchronization with sync.WaitGroup and sync.Mutex.
- Requests:
- - Basic HTTP requests with net/http.
- - Parsing responses (io.Reader).
- Data Scraping:
- - Using colly for web scraping.
- - HTML parsing with goquery.

### 12. Error Handling (Try Except)
- Error Handling Idioms:
- - if err != nil.
- - Wrapping errors with fmt.Errorf.
- Custom Errors:
- - Defining custom error types.
- Panic and Recover:
- - Handling unexpected errors gracefully.

### 13. File Handling and IO
- File Operations:
- - Reading and writing files (os and io/ioutil).
- Buffering:
- - Using bufio for efficient IO.
- File Manipulation:
- - Checking existence, renaming, and deleting files.
