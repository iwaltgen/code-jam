package countingsheep

import "bytes"

func ExampleSolution() {
	in := `5
0
1
2
11
1692`
	solveProblem(bytes.NewReader([]byte(in)))
	// Output:
	// Case #1: INSOMNIA
	// Case #2: 10
	// Case #3: 90
	// Case #4: 110
	// Case #5: 5076
}
