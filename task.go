package main

type Task struct {
	Paragraph string
	Options   Listable
}

func CreateTask(para string, options Listable) Task {
	return Task{Paragraph: para, Options: options}
}
