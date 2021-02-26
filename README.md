[![Go Reference](https://pkg.go.dev/badge/github.com/Pantani/errors.svg)](https://pkg.go.dev/github.com/Pantani/errors)
[![codecov](https://codecov.io/gh/Pantani/errors/branch/master/graph/badge.svg?token=F6HFPDPW9J)](https://codecov.io/gh/Pantani/errors)

# Simple error package

Simple abstraction for errors.

Use this package for create a new error.
An error in Go is any implementing interface with an Error() string method. We overwrite the error object by our error struct:

```go
type Error struct {
	Err   error
	Type  Type
	meta  map[string]interface{}
	stack []string
}
```

To be easier the error construction, the package provides a function named E, which is short and easy to type:

```go
func E(args ...interface{}) *Error
```

E.g.:
- just error:
```go
errors.E(err)
```

- error with message:
```go
errors.E(err, "new message to append")
```

- error with type:
```go
errors.E(err, errors.TypePlatformReques)
```

- error with type and message:
```go
errors.E(err, errors.TypePlatformReques, "new message to append")
```


- error with type and meta:
```go
errors.E(err, errors.TypePlatformRequest, errors.Params{
			"coin":   "Ethereum",
			"method": "CurrentBlockNumber",
		})
```

- error with meta:
```go
errors.E(err, errors.Params{
			"coin":   "Ethereum",
			"method": "CurrentBlockNumber",
		})
```

- error with type and meta:
```go
errors.E(err, errors.TypePlatformRequest, errors.Params{
			"coin":   "Ethereum",
			"method": "CurrentBlockNumber",
		})
```

- error with type, message and meta:
```go
errors.E(err, errors.TypePlatformRequest, "new message to append", errors.Params{
			"coin":   "Ethereum",
			"method": "CurrentBlockNumber",
		})
```


- You can send the errors to sentry using `.PushToSentry()`
```go
errors.E(err, errors.TypePlatformReques).PushToSentry()
```


