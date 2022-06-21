package behavioral

type department interface {
	execute(*patient)
	setNext(department)
}
