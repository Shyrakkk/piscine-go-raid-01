package piscine

import "github.com/01-edu/z01"

func PrintWidthC(x int, last int) {
	for i := 0; i < x; i++ {
		if i == 0 {
			if last == 1 {
				z01.PrintRune('C')
				continue
			}
			z01.PrintRune('A')
		} else if i == x-1 {
			if last == 1 {
				z01.PrintRune('C')
				continue
			}
			z01.PrintRune('A')
		} else {
			z01.PrintRune('B')
		}
	}
	z01.PrintRune(10)
}

func Raid1c(x, y int) {

	if x == 0 && y == 0 {
		return
	}

	if x < 0 {
		return
	}

	if y < 0 {
		return
	}

	PrintWidthC(x, 0)

	if y > 1 {
		for i2 := 1; i2 < y; i2++ {
			if i2 == y-1 {
				PrintWidthC(x, 1)
				break
			} else if x > 1 {
				z01.PrintRune('B')
				PrintSpace(x - 2)
				z01.PrintRune('B')
			} else {
				z01.PrintRune('B')
			}
			z01.PrintRune(10)
		}
	}

}
