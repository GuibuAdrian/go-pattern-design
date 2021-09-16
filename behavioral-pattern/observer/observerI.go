package main

type ObserverI interface {
	update(string)
	getID() string
}
