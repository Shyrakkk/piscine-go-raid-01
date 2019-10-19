package piscine

import "github.com/01-edu/z01"

func PrintSpace(n int) {
	for i := 0; i < n; i++ {
		z01.PrintRune(' ')
	}
}

func PrintWidth(x int) {
	for i := 0; i < x; i++ {
		if i == 0 || i == x-1 {
			z01.PrintRune('o')
		} else {
			z01.PrintRune('-')
		}
	}
	z01.PrintRune(10)
}

func Raid1a(x, y int) {

	if x < 0 {
		x = x * (-1)
	}

	if y < 0 {
		y = x * (-1)
	}

	PrintWidth(x)

	if y > 1 {
		for i2 := 1; i2 < y; i2++ {
			if i2 == y-1 {
				PrintWidth(x)
				break
			} else if x > 1 {
				z01.PrintRune('|')
				PrintSpace(x - 2)
				z01.PrintRune('|')
			} else {
				z01.PrintRune('|')
			}
			z01.PrintRune(10)
		}
	}

}
