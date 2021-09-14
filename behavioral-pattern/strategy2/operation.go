package main

type operation struct {
	op operationI
}

func initOperation(op2 operationI)  operation{
	return operation{
		op: op2,
	}
}

func (op1 operation) setOp(op2 operationI)  {
	op1.op = op2
}