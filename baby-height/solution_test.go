package babyheight

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		in  *input
		out string
	}{
		{
			in: &input{
				index:  1,
				gender: "B",
				mother: newHeight(5, 11),
				father: newHeight(6, 2),
				// 12'1 +5
				// 12'6
				// 6'3 +-4
			},
			out: "Case #1: 5'11\" to 6'7\"",
		},
		{
			in: &input{
				index:  2,
				gender: "G",
				mother: newHeight(5, 11),
				father: newHeight(6, 2),
				// 12'1 -5
				// 11'8
				// 5'10 +-4
			},
			out: "Case #2: 5'6\" to 6'2\"",
		},
		{
			in: &input{
				index:  3,
				gender: "B",
				mother: newHeight(3, 4),
				father: newHeight(3, 4),
				// 6'8 +5
				// 7'1
				// 3'6.5 +-4
			},
			out: "Case #3: 3'3\" to 3'10\"",
		},
		{
			in: &input{
				index:  4,
				gender: "G",
				mother: newHeight(1, 1),
				father: newHeight(1, 0),
				// 2'1 -5
				// 1'8
				// 0'10 +-4
			},
			out: "Case #4: 0'6\" to 1'2\"",
		},
	}

	for _, data := range tests {
		assert.Equal(data.out, compute(data.in).String())
	}
}

func ExampleCompute() {
	in := `4
B 5'11" 6'2"
G 5'11" 6'2"
B 3'4" 3'4"
G 1'1" 1'0"`
	inputs := readInput(bytes.NewReader([]byte(in)))
	for _, input := range inputs {
		fmt.Println(compute(input))
	}
	// Output:
	// Case #1: 5'11" to 6'7"
	// Case #2: 5'6" to 6'2"
	// Case #3: 3'3" to 3'10"
	// Case #4: 0'6" to 1'2"
}
