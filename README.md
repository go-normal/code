## golang error code

![Go](https://github.com/go-normal/code/workflows/Go/badge.svg)

### 1.usage

#### defined error code
```
import "github.com/go-normal/code"

var (
    ErrNotFound = code.MustAdd(100404, "not found")
    ErrInvalidParam := code.MustAdd(1234, "param invalid: %s - %s")
)

```

#### get the error message

```
......

err := doSometing() // return ErrNotFound

if err != nil {
    fmt.Print(err.Message())
}


err = doSometing2() // return ErrInvalidParam

if err != nil {
    fmt.Print(err.Message("mobile", "length != 11"))
}

......


```