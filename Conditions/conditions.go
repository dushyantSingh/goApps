package main

func main() {
	x := 10
	if x > 5 {
		println("X is bigger than 5")
	}
	if x > 100 {
		println("X is very big")
	} else {
		println("X is not very big")
	}

	if x > 5 && x < 10 {
		println("X is not that big")
	} else {
		println("X is slightly big")
	}
	a, b := 11.0, 20.0

	if frac := a / b; frac > 0.5 {
		println("a is more than half of b")
	}

	switchCondintion()
}

func switchCondintion() {
	a := 10

	switch a {
	case 5:
		println("One")
	case 10:
		println("TWO")
	case 15:
		println("Three")
	default:
		println("Defaults")
	}

	switch {
	case a > 100:
		println("Greater")
	case a < 100:
		println("Lower")
	}
}
