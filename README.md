<h1 align="center"><code>as</code></h1>

<p align="center">Convert any value to a known type.</p>

```golang
query := r.URL.Query()

as.Int(query.Get("page"))
as.Float(query.Get("min"))
as.Bool(query.Get("in_stock"))
```

## Installation

```bash
go get github.com/AnatoleLucet/as
```

## Usage

```golang
package main

import (
    "github.com/AnatoleLucet/as"
)

func main() {
    as.String(123)              // "123"
    as.String(1.234)            // "1.234"
    as.String(true)             // "true"
    as.String([]byte("hello"))  // "hello"
    as.String([]int{1,2,3})     // "[1 2 3]"

    as.Int("123")               // 123
    as.Int(1.234)               // 1
    as.Int(true)                // true
    as.Int([]byte("123"))       // 123

    as.Bool("true")             // true
    as.Bool("yes")              // true
    as.Bool(1)                  // true
    as.Bool("false")            // false
    as.Bool("no")               // false
    as.Bool(0)                  // false

    as.Slice(as.Int)([]string{"1", "2", "3"})

    // and many more
}
```

#### `as.Value()`

as.Value() is a special generic type, purpose built for cases where you want to abstract an arbitrary value in a known type.

```golang
// With as.Value you can store multiple types in a single slice.
values := []as.V{
    as.Value("hello"),
    as.Value(123),
    as.Value(1.234),
}

str, _ := values[0].String()
num, _ := values[1].Int()
flt, _ := values[2].Float()
```
