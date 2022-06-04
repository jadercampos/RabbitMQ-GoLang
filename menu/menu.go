package menu

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jadercampos/RabbitMQ-GoLang/hello"
	"github.com/jadercampos/RabbitMQ-GoLang/pubSub"
	"github.com/jadercampos/RabbitMQ-GoLang/workQueues"
)

var comando int

func ExibeMenu() {

	fmt.Println("<<< Exemplos com RabbitMQ >>>")
	fmt.Println("***************************************************************")
	fmt.Println("\nOpções: ")

	fmt.Println("\n> Hello Word:")
	fmt.Println("1- Sender")
	fmt.Println("2- Receiver")

	fmt.Println("\n> Work Queues:")
	fmt.Println("3- New Task")
	fmt.Println("4- Worker")

	fmt.Println("\n> Pub/Sub:")
	fmt.Println("5- Emit Log")
	fmt.Println("6- Receive Log")

	fmt.Println("\n0- Sair do Programa")

	LeComando()
}
func LeComando() {
	fmt.Print("Digite a opção e aperte Enter: ")
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
	switch comando {
	case 1:
		hello.Sender("hello", "Oi mundo!")
	case 2:
		hello.Receiver("hello")
	case 3:
		workQueues.PublicaORole()
	case 4:
		workQueues.ConsomeORole()
	case 5:
		pubSub.EmitLog()
	case 6:
		pubSub.ReceiveLog()
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		VoltarAoMenu()
	}

}

func LimpaConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func VoltarAoMenu() {
	fmt.Print("Aperte 1 para voltar ao menu ou qualquer outra tecla para sair: ")
	fmt.Scan(&comando)
	if comando == 1 {
		ExibeMenu()
	} else {
		os.Exit(0)
	}
}
