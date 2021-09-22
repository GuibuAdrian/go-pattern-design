package main

import "fmt"

type Developer struct {
	name string
	salary float64
}

func initDeveloper(name string, salary float64) *Developer {
	return &Developer{
		name: name,
		salary: salary,
	}
}

func (dev *Developer) add (emp EmployeeI)  {
	//This is leaf node
}

func (dev *Developer) getPos(i int) EmployeeI {
	//This is leaf node
	return nil
}

func (dev *Developer) getName() string {
	return dev.name
}

func (dev *Developer) getSalary() float64 {
	return dev.salary
}

func (dev *Developer) remove(emp EmployeeI)  {
	//This is leaf node
}

func (dev *Developer) print() {
	fmt.Println("\t\t\tName", dev.getName())
	fmt.Println("\t\t\tSalary:", dev.getSalary())
}