```go
func main() {
    var x int = 10
    var ptr *int = &x
    fmt.Println(*ptr) // Output: 10

    *ptr = 20
    fmt.Println(x) // Output: 20, x is modified through the pointer

    ptr = nil // Set the pointer to nil
    fmt.Println(*ptr) // This will cause a runtime panic: runtime error: invalid memory address or nil pointer dereference
}
```