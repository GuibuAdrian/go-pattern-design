package structural

type EmployeeI interface {
	add(emp EmployeeI)
	remove(emp EmployeeI)
	getPos(pos int) EmployeeI
	getName() string
	getSalary() float64
	print()
}
