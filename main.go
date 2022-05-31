package main

import "github.com/jadercampos/RabbitMQ-GoLang/hello"

func main() {

}

func HelloWorld() {
	hello.Sender("hello", "Oi pessoal")
	hello.Receiver("hello")
}
