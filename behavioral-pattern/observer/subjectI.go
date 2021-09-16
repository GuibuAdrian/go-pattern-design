package main

type subjectI interface {
	register(ObserverI)
	deregister(ObserverI)
	notifyAll()
}