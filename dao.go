package main

type ConnectionPort struct {
	id   int
	name string
}

type ConnectionClient struct {
	id    int
	name  string
	ports []ConnectionPort
}
