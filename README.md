# error

[![Build Status](https://secure.travis-ci.org/goforgery/error.png?branch=master)](http://travis-ci.org/goforgery/error)

Simple error handler for Forgery2.

## Use

When running in the environment `prod` any calls to `panic()` will be caught and a 500 response will be returned. You can set the title of the 500 response by passing a string as the first argument to create.

```javascript
package main

import (
	"github.com/goforgery/forgery2"
	"github.com/goforgery/error"
)

func main() {
	app := f.CreateApp()
	app.Set("env", "prod")
	app.Use(error.Create("Something went wrong!"))
	app.Get("/", func(req *f.Request, res *f.Response, next func()) {
		panic("panic")
	})
	app.Listen(3000)
}
```

## Test

    go test
