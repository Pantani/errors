# Simple log package

Simple abstraction for errors.

Use this package for create a new error.
An error in Go is any implementing interface with an Error() string method. We overwrite the error object by our error struct:

```
type Error struct {
	Err   error
	Type  Type
	meta  map[string]interface{}
	stack []string
}
```

To be easier the error construction, the package provides a function named E, which is short and easy to type:

`func E(args ...interface{}) *Error`

E.g.:
- just error:
`errors.E(err)`

- error with message:
`errors.E(err, "new message to append")`

- error with type:
`errors.E(err, errors.TypePlatformReques)`

- error with type and message:
`errors.E(err, errors.TypePlatformReques, "new message to append")`

- error with type and meta:
```
errors.E(err, errors.TypePlatformRequest, errors.Params{
			"coin":   "Ethereum",
			"method": "CurrentBlockNumber",
		})
```

- error with meta:
```
errors.E(err, errors.Params{
			"coin":   "Ethereum",
			"method": "CurrentBlockNumber",
		})
```

- error with type and meta:
```
errors.E(err, errors.TypePlatformRequest, errors.Params{
			"coin":   "Ethereum",
			"method": "CurrentBlockNumber",
		})
```

- error with type, message and meta:
```
errors.E(err, errors.TypePlatformRequest, "new message to append", errors.Params{
			"coin":   "Ethereum",
			"method": "CurrentBlockNumber",
		})
```


- You can send the errors to sentry using `.PushToSentry()`
`errors.E(err, errors.TypePlatformReques).PushToSentry()`

