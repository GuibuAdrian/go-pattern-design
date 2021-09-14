package main

type operationI interface {
	doOperation(a int, b int) int
}
