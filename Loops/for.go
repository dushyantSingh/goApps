package main

func main() {
	for i := 0; i < 4; i++ {
		println("Loop index", i)
	}
	println("-------")

	for i := 0; i < 4; i++ {
		if i > 2 {
			break
		}
		println("Loop index", i)
	}

	println("-------")

	for i := 0; i < 4; i++ {
		if i > 2 {
			continue
		}
		println("Loop index", i)
	}

	println("-------")
	b := 0
	for {
		if b > 4 {
			break
		}
		b++
		println("Foprrrrr ", b)
	}
}
