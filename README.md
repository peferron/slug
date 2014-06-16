# slug [![GoDoc](https://godoc.org/github.com/peferron/slug?status.png)](https://godoc.org/github.com/peferron/slug) [![Build Status](https://drone.io/github.com/peferron/slug/status.png)](https://drone.io/github.com/peferron/slug/latest) [![Coverage Status](https://coveralls.io/repos/peferron/slug/badge.png?branch=master)](https://coveralls.io/r/peferron/slug?branch=master)

slug is a Go library for generating URL-friendly slugs. Examples:
- A Lovely Day → a-lovely-day
- 北京 → bei-jing

#### Usage

```go
package main

import "github.com/peferron/slug"

func main() {
    s := slug.Slug("Hello World!", 50, "-")
    print(s) // Prints "hello-world"
}
```

[API reference](https://godoc.org/github.com/peferron/slug)
