# Unexpected Memory Address Changes in Go

This repository demonstrates a common source of confusion in Go: how the memory address of a variable can change unexpectedly.  While Go's garbage collection handles memory allocation automatically, understanding how variables are stored in memory is crucial for writing robust and efficient code.

The example shows how printing the memory address of a variable using the `&` operator reveals different addresses for variables declared in different scopes or with different initialization methods.