package main

import (
	"github.com/jadercampos/RabbitMQ-GoLang/hello"
	"github.com/jadercampos/RabbitMQ-GoLang/workQueues"
)

func main() {
	workQ()
}

func helloWorld() {
	hello.Sender("hello", "Oi pessoal")
	hello.Receiver("hello")
}

func workQ() {
	workQueues.PublicaORole()
	//workQueues.ConsomeORole()
}
