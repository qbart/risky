# risky go library

Risky parses strings into corresponding types but ignores errors and default or zero values are returned instead.

Bit unsafe but gives values when you need them or you can trust the input.

## Usage

```go
type Animal struct {
    Name string
}

riksy.JSON[Animal](`{"name":"kiwi"}`) // returns an instance of struct{Name: "kiwi"}
riksy.JSON[Animal](`abc`) // returns zero value of struct{Name: ""}
risky.ParseInt("123") // returns 123
risky.ParseInt("invalid") // returns 0
```