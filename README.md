# go-april
Golang App.

[See numbers.go](numbers.go)

The numbers are handled by this variable:
```go
type elements []int
```

This function is attached to ```elements```
```go
func (e elements) print() {
	for _, item := range e {
		fmt.Println(item)
	}

}
```