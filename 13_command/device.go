package main

type Device interface {
	Open()
	Close()
	Restart()
}
