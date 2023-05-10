package main

type Server interface {
	HandlerReq(string, string) (int, string)
}
