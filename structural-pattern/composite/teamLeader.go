package main

import "fmt"

type TeamLeader struct {
	name string
	salary float64
	team []EmployeeI
}

func initTeamLeader(name string, salary float64) *Manager{
	employees := make([]EmployeeI,1)
	return &Manager{
		name: name,
		salary: salary,
		employees: employees,
	}
}

func (tmL *TeamLeader) getPos(pos int) EmployeeI {
	return tmL.team[pos]
}

func (tmL *TeamLeader) add(emp EmployeeI)  {
	tmL.team = append(tmL.team, emp)
}

func (tmL *TeamLeader) remove(emp EmployeeI)  {
	tmL.team = removeFromSlice(tmL.team, emp)
}

func (tmL *TeamLeader) getName() string {
	return tmL.name
}

func (tmL *TeamLeader) getSalary() float64 {
	return tmL.salary
}

func (tmL *TeamLeader) print() {
	fmt.Println("\tTeamLeader:")
	fmt.Println("\tName:", tmL.getName())
	fmt.Println("\tSalary:", tmL.getSalary())

	fmt.Println("\t\tEmployees:")
	for _, emp := range tmL.team {
		emp.print()
	}
	fmt.Println()
}