# slug [![GoDoc](https://godoc.org/github.com/peferron/slug?status.png)](https://godoc.org/github.com/peferron/slug) [![Build Status](https://drone.io/github.com/peferron/slug/status.png)](https://drone.io/github.com/peferron/slug/latest) [![Coverage Status](https://coveralls.io/repos/peferron/slug/badge.png?branch=master)](https://coveralls.io/r/peferron/slug?branch=master)

slug is a Go library for generating URL-friendly [slugs](http://en.wikipedia.org/wiki/Slug_%28web_publishing%29#Slug).

#### Examples

- A Lovely Day → a-lovely-day
- Living in 北京 → living-in-bei-jing

#### Usage

```go
package main

import "github.com/peferron/slug"

func main() {
    s := slug.Generate("Hello World")
    print(s) // Prints "hello-world"
}
```

[API reference](https://godoc.org/github.com/peferron/slug)
