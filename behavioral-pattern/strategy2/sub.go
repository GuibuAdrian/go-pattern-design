package main

type sub struct {

}

func (op sub) doOperation(a, b int) int {
	return a - b
}
