package main

var _x = 1

var (
	_y = 2
	_z = 3
)

func main_var() {
	var (
		a string = "hi"
	)

	var n string

	n = a

	println(n, _x, _y, _z)
}
