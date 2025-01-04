package datastructures_test

import (
	"testing"

	ds "github.com/Ajstraight619/gogit/data_structures"
)

func TestStackPush(t *testing.T) {
	stack := ds.Stack[int]{}
	stack.Push(1)

}
