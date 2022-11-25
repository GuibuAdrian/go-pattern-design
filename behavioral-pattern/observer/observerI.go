package behavioral

type ObserverI interface {
	update(string)
	getID() string
}
