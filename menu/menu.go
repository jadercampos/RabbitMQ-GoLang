package menu

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/jadercampos/RabbitMQ-GoLang/hello"
	"github.com/jadercampos/RabbitMQ-GoLang/pubSub"
	"github.com/jadercampos/RabbitMQ-GoLang/routing"
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

	fmt.Println("\n> Routing:")
	fmt.Println("7- Emita")
	fmt.Println("8- Receba")

	fmt.Println("\n0- Sair do Programa")

	LeComando()
}
func LeComando() {
	fmt.Print("\nDigite a opção e aperte Enter: ")
	fmt.Scan(&comando)
	fmt.Println("\nO comando escolhido foi", comando)
	switch comando {
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
			severities, valido = scanUserInput("\nDigite os tipos de logs que deseja obter separados por espaço [info|warning|error|*]: ", []string{"info", "warning", "error"})
		}
		fmt.Println("\nExibindo logs do tipo: ", severities, "\n")
		routing.RecebaORole("logs_direct", "direct", severities)
	case 0:
		fmt.Println("\nSaindo do programa")
		os.Exit(0)
	default:
		fmt.Println("\nNão conheço este comando")
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

func scanUserInput(msg string, validValues []string) ([]string, bool) {
	fmt.Print(msg)
	scanner := bufio.NewScanner(os.Stdin)
	var informedValues = validValues
	valido := true
	for scanner.Scan() {
		digitado := scanner.Text()
		if digitado != "" && digitado != "\n" {
			informedValues = strings.Fields(digitado)
			if digitado == "*" {
				informedValues = validValues
			}
			if !temAlgum(validValues, informedValues) {
				fmt.Println("\nValor informado é inválido: ", informedValues)
				valido = false
			}
			break
		}
	}
	return informedValues, valido
}

func temAlgum(validValues []string, informedValues []string) bool {
	var temAlgum bool
	for _, item := range informedValues {
		i := sort.SearchStrings(validValues, item)
		tem := i < len(informedValues) && informedValues[i] == item
		if tem {
			temAlgum = true
		}
	}
	return temAlgum
}
