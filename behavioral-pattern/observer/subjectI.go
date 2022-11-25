package behavioral

type subjectI interface {
	register(ObserverI)
	deregister(ObserverI)
	notifyAll()
}
