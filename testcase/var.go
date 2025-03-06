package main

var x = 1

var (
	y = 2
	z = 3
)

func main() {
	var (
		a string = "hi"
	)

	var n string

	n = a

	println(n, x, y, z)
}
