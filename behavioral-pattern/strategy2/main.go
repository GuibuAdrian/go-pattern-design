package main

import (
	"fmt"
)
func main() {
	operation := initOperation(sum{})
	fmt.Println(operation.op.doOperation(4, 5))

	operation = initOperation(sub{})
	fmt.Println(operation.op.doOperation(2,3))

}
