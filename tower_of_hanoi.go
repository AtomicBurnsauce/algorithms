/*
	Tower of Hanoi, is a mathematical puzzle which consists of three towers
	(pegs) and more than one ring. The puzzle starts with the disks a neat
	stack in ascending order of size on one rod, the smallest on top, thus
	making a conical shape.
*/

package main

import "fmt"

type solver interface {
	play(int)
}

type towers struct {
	// an empty struct
}

// play is sole method requried to implement solver type
func (t *towers) play(n int) {
	t.moveN(n, 1, 2, 3)
}

// recursive algorithm
func (t *towers) moveN(n, from, to, via int) {
	if n > 0 {
		t.moveN(n-1, from, via, to)
		t.moveM(from, to)
		t.moveN(n-1, via, to, from)
	}
}

func (t *towers) moveM(from, to int) {
	fmt.Println("Move disk from rod", from, "to rod", to)
}

func main() {
	var t solver
	t =new(towers) // type towers must satisfy solver interface
	t.play(4)
}