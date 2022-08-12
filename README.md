# Convert ANY to ANY type in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/dpanic/conv)](https://goreportcard.com/report/github.com/dpanic/any)

Converting String to []string, or string to float64 could be handy sometimes, when your project has inconsistent input.

## Usage
``` Go
package main

import (
	"fmt"

	"github.com/dpanic/convert"
)

func main() {
	res := convert.ToInt("123")
	fmt.Printf("%+v %t\n", res, res)
}
```

Go Playground is here https://go.dev/play/p/Z5phbxlWGg0


## Features
* Convert basic types to any other basic type
* Same as above but with * pointer

### Primitive type conversions:
* ToBoolP -> *bool
* ToIntP -> *int
* ToInt64P -> *int64
* ToFloat32P -> *float32
* ToFloat64P -> *float64
* ToStringP -> *string
* ToString -> string
* ToInt -> int
* ToInt64 -> int64
* ToFloat -> float64
* ToBool -> bool

### Composite type conversions:
* ToSliceOfString -> []string
* ToSliceOfFloat -> []float64
* ToMapOfStrings -> map[string]string
* ToMapOfInterfaces -> map[string]interface{}
* ToSliceOfMap -> []map[string]string
* ToSliceOfMapOfInterfaces -> []map[string]interface{}
* ToSliceOfBool -> []bool
