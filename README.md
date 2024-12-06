# Go: Nil Pointer Dereference

This repository demonstrates a common error in Go: nil pointer dereference.

## Bug Description
The code attempts to dereference a nil pointer, leading to a runtime panic. This is a classic issue that can be difficult to track down if not handled carefully.

## How to Reproduce
1. Clone the repository.
2. Run the `bug.go` file.
3. Observe the runtime panic.

## Solution
The `bugSolution.go` file provides a corrected version. Always check for `nil` before dereferencing a pointer.