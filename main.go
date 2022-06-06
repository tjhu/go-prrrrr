// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Stream interface {
	Exec()
	Done()
}

type Operator struct {
}

func (op *Operator) Exec() {
	fmt.Println("Operator Exec")
}

func RunDAG(s Stream) {
	s.Exec()
}

func (op *Operator) Done() {
	RunDAG(op)
}

type Source struct {
	Operator
}

func (src *Source) Exec() {
	fmt.Println("Source Exec")
}

func Bar(s Stream) {
	s.Exec()
}

func main() {
	//op := Operator{}
	src := Source{}

	var srcs Stream = &src

	src.Exec()
	srcs.Exec()
	Bar(srcs)

	src.Done()
	srcs.Done()

}
