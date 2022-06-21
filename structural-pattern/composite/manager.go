package structural

import "fmt"

type Manager struct {
	name      string
	salary    float64
	employees []EmployeeI
}

func initManager(name string, salary float64) *Manager {
	employees := make([]EmployeeI, 1)
	return &Manager{
		name:      name,
		salary:    salary,
		employees: employees,
	}
}

func (man *Manager) getPos(pos int) EmployeeI {
	return man.employees[pos]
}

func (man *Manager) add(emp EmployeeI) {
	man.employees = append(man.employees, emp)
}

func (man *Manager) remove(emp EmployeeI) {
	man.employees = removeFromSlice(man.employees, emp)
}

func removeFromSlice(employees []EmployeeI, empToRemove EmployeeI) []EmployeeI {
	employeesLength := len(employees)
	for i, employee := range employees {
		if empToRemove.getName() == employee.getName() {
			employees[employeesLength-1], employees[i] = employees[i], employees[employeesLength-1]
			return employees[:employeesLength-1]
		}
	}

	return employees
}

func (man *Manager) getName() string {
	return man.name
}

func (man *Manager) getSalary() float64 {
	return man.salary
}

func (man *Manager) print() {
	fmt.Println("Manager:")
	fmt.Println("Name:", man.getName())
	fmt.Println("Salary:", man.getSalary())

	fmt.Println("Employees:")
	for _, emp := range man.employees {
		emp.print()
	}
}
