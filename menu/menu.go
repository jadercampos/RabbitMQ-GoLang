package menu

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jadercampos/RabbitMQ-GoLang/hello"
	"github.com/jadercampos/RabbitMQ-GoLang/pubSub"
	"github.com/jadercampos/RabbitMQ-GoLang/rabbitUtils"
	"github.com/jadercampos/RabbitMQ-GoLang/routing"
	"github.com/jadercampos/RabbitMQ-GoLang/workQueues"
)

var selectedCommand int

func ShowMenu() {

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

	fmt.Println("\n> Routing:")
	fmt.Println("7- Emita")
	fmt.Println("8- Receba")

	fmt.Println("\n0- Sair do Programa")

	ReadCommand()
}

func ReadCommand() {
	fmt.Print("\nDigite a opção e aperte Enter: ")
	fmt.Scan(&selectedCommand)
	fmt.Println("\nO comando escolhido foi", selectedCommand)
	switch selectedCommand {
	case 1:
		hello.Sender("hello", "Oi mundo!")
	case 2:
		hello.Receiver("hello")
	case 3:
		workQueues.PublicaORole("task_queue")
	case 4:
		workQueues.ConsomeORole("task_queue")
	case 5:
		pubSub.EmitLog("logs", "fanout")
	case 6:
		pubSub.ReceiveLog("logs", "fanout")
	case 7:
		routing.EmitaORole("logs_direct", "direct")
	case 8:
		var severities []string
		var valido bool
		for valido == false {
			severities, valido = rabbitUtils.ScanUserInput("\nDigite os tipos de logs que deseja obter separados por espaço [info|warning|error|*]: ", []string{"info", "warning", "error"})
		}
		fmt.Println("\nExibindo logs do tipo: ", severities, "\n")
		routing.RecebaORole("logs_direct", "direct", severities)
	case 0:
		fmt.Println("\nSaindo do programa")
		os.Exit(0)
	default:
		fmt.Println("\nNão conheço este comando")
		BackToMenu()
	}
}

func CleanConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func BackToMenu() {
	fmt.Print("Aperte 1 para voltar ao menu ou qualquer outra tecla para sair: ")
	fmt.Scan(&selectedCommand)
	if selectedCommand == 1 {
		ShowMenu()
	} else {
		os.Exit(0)
	}
}
