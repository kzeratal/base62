# Base62

[![Build Status](https://app.travis-ci.com/kzeratal/base62.svg?branch=main)](https://app.travis-ci.com/kzeratal/base62) [![Coverage Status](https://coveralls.io/repos/github/kzeratal/base62/badge.svg?branch=main)](https://coveralls.io/github/kzeratal/base62?branch=main)

Base62 encoding written in Go.

```go
package main

import (
	"fmt"

	"github.com/kzeratal/base62"
)

func main() {
	str := base62.Encode(89569285645)
	fmt.Println(str) // "1ZlfarV"
}
```
