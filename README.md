# Golang metric system

This package provides basic metric system conversions.
To start, in your project, run `go get github.com/dragonDScript/golang-metric-system`

Now, in your project, you can call all this package's methods.

```go
package main

import (
    "fmt"
    "github.com/dragonDScript/golang-metric-system"
)

func main() {
    // Convert 10 meters to decimetres
    conversion := metric.Convert(Metre * 10, Decimetre)
    fmt.Println(conversion)
    fmt.Println(metres.Metre, metres.Decimetre)
}

```
