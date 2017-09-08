/*
Package babyheight Baby Height

	* https://code.google.com/codejam/contest/2334486/dashboard#s=p1
	* https://www.acmicpc.net/problem/12353
*/
package babyheight

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type input struct {
	index  int
	gender string
	mother height
	father height
}

type output struct {
	index int
	min   height
	max   height
}

func (o *output) String() string {
	return fmt.Sprintf("Case #%d: %v to %v", o.index, o.min, o.max)
}

const (
	maxFeet = 10
	maxInch = 12
)

type height struct {
	feet int
	inch int
}

func newHeight(feet, inch int) height {
	return height{
		feet: feet,
		inch: inch,
	}
}

func parseHeight(text string) height {
	var feet, inch int
	fmt.Sscanf(text, "%d'%d\"", &feet, &inch)
	return height{feet, inch}
}

func (h height) add(o height) height {
	feet := h.feet + o.feet
	inch := h.inch + o.inch
	return newHeight(feet+(inch/maxInch), inch%maxInch)
}

func (h height) sub(o height) height {
	feet := h.feet - o.feet
	inch := h.inch - o.inch
	if inch < 0 {
		return newHeight(feet-1, maxInch+inch)
	}
	return newHeight(feet, inch)
}

func (h height) div(n int) (height, bool) {
	feet := h.feet / n
	modFeet := h.feet % n
	if modFeet != 0 {
		h.inch += modFeet * maxInch
	}
	inch := int(float64(h.inch)/float64(n) + 0.5)
	return newHeight(feet, inch), h.inch%n != 0
}

func (h height) String() string {
	return fmt.Sprintf("%d'%d\"", h.feet, h.inch)
}

func compute(in *input) *output {
	std := in.father.add(in.mother)
	switch in.gender {
	case "B":
		std = std.add(newHeight(0, 5))
	case "G":
		std = std.sub(newHeight(0, 5))
	}

	std, rounded := std.div(2)
	if rounded {
		return &output{
			index: in.index,
			min:   std.sub(newHeight(0, 4)),
			max:   std.add(newHeight(0, 4-1)),
		}
	}
	return &output{
		index: in.index,
		min:   std.sub(newHeight(0, 4)),
		max:   std.add(newHeight(0, 4)),
	}
}

func readInput(in io.Reader) []*input {
	var n int
	fmt.Fscanf(in, "%d\n", &n)
	reader := bufio.NewReader(in)
	result := make([]*input, 0, n)
	for i := 0; i < n; i++ {
		text, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		tokens := strings.Split(text, " ")
		result = append(result, &input{
			index:  i + 1,
			gender: tokens[0],
			mother: parseHeight(tokens[1]),
			father: parseHeight(tokens[2]),
		})
	}
	return result
}
