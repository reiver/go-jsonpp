# go-jsonpp

Package **jsonpp** provides tools for **pretty printing JSON**.

**jsonpp** = **json pretty print**

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-jsonpp

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-jsonpp?status.svg)](https://godoc.org/sourcecode.social/reiver/go-jsonpp)

## Example

Here is an example usage:

```go
import "sourcecode.social/reiver/go-jsonpp"

// ...


var jsn []byte // <--- contains JSON

// ...

jsonpp.PrettyPrint(jsn) // outputs the pretty-printed JSON to the STDOUT.
```

Can also send it to an `io.Writer`

```go
import "sourcecode.social/reiver/go-jsonpp"

// ...


var jsn []byte // <--- contains JSON

// ...

jsonpp.FPrettyPrint(writer, jsn) // outputs the pretty-printed JSON to an io.Writer.
```

Or return it as a `string`:

```go
import "sourcecode.social/reiver/go-jsonpp"

// ...


var jsn []byte // <--- contains JSON

// ...

s := jsonpp.SPrettyPrint(jsn) // returns the pretty-printed JSON as a string.
```


## Import

To import package **jsonpp** use `import` code like the follownig:
```
import "sourcecode.social/reiver/go-jsonpp"
```

## Installation

To install package **jsonpp** do the following:
```
GOPROXY=direct go get https://sourcecode.social/reiver/go-jsonpp
```

## Author

Package **jsonpp** was written by [Charles Iliya Krempeaux](http://changelog.ca)
