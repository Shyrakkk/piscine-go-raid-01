package piscine

import "github.com/01-edu/z01"

func PrintWidth2(x int, last int) {
	for i := 0; i < x; i++ {
		if i == 0 {
			if last == 1 {
				z01.PrintRune(92)
				continue
			}
			z01.PrintRune('/')
		} else if i == x-1 {
			if last == 1 {
				z01.PrintRune('/')
				continue
			}
			z01.PrintRune(92)
		} else {
			z01.PrintRune('*')
		}
	}
	z01.PrintRune(10)
}

func Raid1b(x, y int) {

	if x == 0 && y == 0 {
		return
	}

	if x < 0 {
		return
	}

	if y < 0 {
		return
	}

	PrintWidth2(x, 0)

	if y > 1 {
		for i2 := 1; i2 < y; i2++ {
			if i2 == y-1 {
				PrintWidth2(x, 1)
				break
			} else if x > 1 {
				z01.PrintRune('*')
				PrintSpace(x - 2)
				z01.PrintRune('*')
			} else {
				z01.PrintRune('*')
			}
			z01.PrintRune(10)
		}
	}

}
