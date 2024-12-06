```go
func main() {
    var x int = 10
    var ptr *int = &x
    fmt.Println(*ptr) // Output: 10

    *ptr = 20
    fmt.Println(x) // Output: 20, x is modified through the pointer

    ptr = nil // Set the pointer to nil
    if ptr != nil {
        fmt.Println(*ptr) // This line will not be executed
    } else {
        fmt.Println("Pointer is nil") // Handle the nil pointer gracefully
    }
}
```