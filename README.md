# atomic
[![PkgGoDev](https://pkg.go.dev/badge/github.com/hslam/atomic)](https://pkg.go.dev/github.com/hslam/atomic)
[![Build Status](https://github.com/hslam/atomic/workflows/build/badge.svg)](https://github.com/hslam/atomic/actions)
[![codecov](https://codecov.io/gh/hslam/atomic/branch/master/graph/badge.svg)](https://codecov.io/gh/hslam/atomic)
[![Go Report Card](https://goreportcard.com/badge/github.com/hslam/atomic)](https://goreportcard.com/report/github.com/hslam/atomic)
[![LICENSE](https://img.shields.io/github/license/hslam/atomic.svg?style=flat-square)](https://github.com/hslam/atomic/blob/master/LICENSE)

Package atomic provides low-level atomic memory primitives useful for implementing synchronization algorithms.

## Feature
* Int8
* Int16
* Int32
* Int64
* Uint8
* Uint16
* Uint32
* Uint64
* Uintptr
* Pointer
* Float32
* Float64
* Bool
* String
* Bytes
* Value

## Get started

### Install
```
go get github.com/hslam/atomic
```
### Import
```
import "github.com/hslam/atomic"
```
### Usage
#### Example
```go
package main

import (
	"fmt"
	"github.com/hslam/atomic"
)

func main() {
	str := atomic.NewString("")
	str.Store("Hi")
	str.Swap("Hello")
	str.Add(" atomic")
	str.CompareAndSwap("Hello atomic", "Hello World")
	fmt.Println(str.Load())
}
```

### Output
```
Hello World
```

### License
This package is licensed under a MIT license (Copyright (c) 2020 Meng Huang)

### Author
atomic was written by Meng Huang.
