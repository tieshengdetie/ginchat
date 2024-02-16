package main

func main() {
	switch f() {
	case true:
		println(1)
	case false:
		println(0)
	default:
		println(-1)
	}
}

func f() bool {
	return false
}
