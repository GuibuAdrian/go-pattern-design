package behavioral

import "fmt"

type item struct {
	observerList []ObserverI
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}
func (i *item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
func (i *item) register(o ObserverI) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o ObserverI) {
	i.observerList = removeFromslice(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, Observer := range i.observerList {
		Observer.update(i.name)
	}
}

func removeFromslice(observerList []ObserverI, observerToRemove ObserverI) []ObserverI {
	ObserverListLength := len(observerList)
	for i, Observer := range observerList {
		if observerToRemove.getID() == Observer.getID() {
			observerList[ObserverListLength-1], observerList[i] = observerList[i], observerList[ObserverListLength-1]
			return observerList[:ObserverListLength-1]
		}
	}
	return observerList
}
