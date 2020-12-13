package main

// These are two constants with the same value. However, the first one (2e1)
// does not seem to be handled correctly.
const (
	twenty = 2e1
	//twenty = 20
)

func main() {
	x := uint32(5)
	println(x / (twenty / 1))
}
