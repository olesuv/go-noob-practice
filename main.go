package main

import "fmt"

func main() {
	// test 1: correct
	a := [3]int{1, 2, 3}
	b := a                  // b = [1, 2, 3]
	b[0] = 9                // b = [9, 2, 3]
	fmt.Println("1:", a, b) // 1: [1, 2, 3], [9, 2, 3]

	// test 2: wrong
	base := [5]int{10, 20, 30, 40, 50}
	s := base[1:3]                             // s = [20, 30]
	s[0] = 99                                  // s = [99, 30]
	fmt.Println("2:", base, s, len(s), cap(s)) // 2: [10, 20, 30, 40, 50], [20, 30], 2, 2

	// test 3: wrong
	x := []int{1, 2, 3, 4}
	y := x[:2]                 // y = [1, 2]
	z := append(y, 7)          // z = [1, 2, 7]
	fmt.Println("3:", x, y, z) // 3: [1, 2, 3, 4], [1, 2], [1, 2, 7]

	// test 4: wrong
	p := make([]int, 2, 3)
	p[0], p[1] = 1, 2       // p = [1, 2, _]
	q := append(p, 3)       // q = [1, 2, 3]
	q[0] = 8                // q = [8, 2, 3]
	fmt.Println("4:", p, q) // 4: [1, 2], [8, 2, 3]

	// test 5: wrong
	src := []int{4, 5, 6}
	dst := make([]int, 2)          // dst = [_, _]
	n := copy(dst, src)            // n = [4, 5]
	dst[0] = 99                    // dst = [99]
	fmt.Println("5:", src, dst, n) // 5: [4, 5, 6], [99], [4, 5]

	// retro:
	// | Operation | Copies elements? | Shares backing array? |
	// |---|---|---|
	// | `b := a` (arrays) | Yes | No |
	// | `s := arr[1:3]` | No | Yes |
	// | `s2 := s1` (slices) | No | Yes |
	// | `append(s, x)` | Sometimes, if reallocation is needed | Depends on capacity |
	// | `copy(dst, src)` | Yes, into destination storage | Does not change existing storage relationships |
}
