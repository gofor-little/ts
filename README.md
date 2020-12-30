## A package for common thread safe data structures

![GitHub tag (latest SemVer pre-release)](https://img.shields.io/github/v/tag/gofor-little/ts?include_prereleases)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gofor-little/ts)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/gofor-little/ts/main/LICENSE)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/gofor-little/ts/Go)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofor-little/ts)](https://goreportcard.com/report/github.com/gofor-little/ts)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gofor-little/ts)](https://pkg.go.dev/github.com/gofor-little/ts)

### Introduction
* Currently supports linked lists and slices
* No dependencies outside the standard library

### Example
```go
package main

import (
	"fmt"

	"github.com/gofor-little/ts"
)

func main() {
    list := ts.LinkedList{}
    list.Push("Some string")

    item := list.GetTail()
    fmt.Println(item)

    item = list.Pop()
    fmt.Println(item)

    slice := ts.Slice{}

    slice.Add(1, 2, 3)
    fmt.Println(slice.Elements)

    slice.Remove(1, 2, 3)
    fmt.Println(slice.Elements)
}
```

### Testing
Run ```go test ./...``` in the root directory.