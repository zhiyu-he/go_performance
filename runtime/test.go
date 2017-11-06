package main

//  go build -gcflags "-N -l" -o test test.go
// go tool compile -help
// -N disable optimizations
// -l disable inlining
func main() {
	s := "hello world"
	println(s)
}
