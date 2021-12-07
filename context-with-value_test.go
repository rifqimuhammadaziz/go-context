package gocontext

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(t *testing.T) {
	// Parent
	contextA := context.Background()

	/**
	Child of contextA
	context.WithValue(parent, key, value)
	*/
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	// Child of contextB
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	// Child of contextC
	contextF := context.WithValue(contextC, "f", "F")

	// Child of contextF
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println("----- Print Value -----")
	fmt.Println(contextF.Value("f")) // contextF has value 'f'
	fmt.Println(contextF.Value("c")) // contextF has value 'c' (from parent)
	fmt.Println(contextF.Value("b")) // contextF has no parent 'b'
}
