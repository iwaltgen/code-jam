/*
Package countingsheep Counting Sheep

	* https://code.google.com/codejam/contest/6254486/dashboard
	* https://www.acmicpc.net/problem/14381
*/
package countingsheep

import (
	"fmt"
	"io"
	"math"
)

func compute(index, n int) {
	if n == 0 {
		fmt.Printf("Case #%d: %s\n", index, "INSOMNIA")
		return
	}

	nmap := map[int]bool{}
	for i := 1; i < math.MaxInt32; i++ {
		number := i * n
		nn := number
		for {
			mn := nn % 10
			nn = nn / 10

			if !nmap[mn] {
				nmap[mn] = true
				if 10 <= len(nmap) {
					fmt.Printf("Case #%d: %d\n", index, number)
					return
				}
			}

			if nn == 0 {
				break
			}
		}
	}
	fmt.Printf("Case #%d: %s\n", index, "INSOMNIA")
}

func solveProblem(in io.Reader) {
	var count, n int
	fmt.Fscanf(in, "%d\n", &count)
	for i := 0; i < count; i++ {
		fmt.Fscanf(in, "%d\n", &n)
		compute(i+1, n)
	}
}
