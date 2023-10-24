package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Patient struct {
	MRnumber string
	Gender   string
}

type HospitalQueue struct {
	Queue []Patient
}

func (h *HospitalQueue) HandleIN() {
	fmt.Println("handle in")
}

func main() {
	fmt.Println(" ===[ Hospital Queue Aplication ]===")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		command := strings.Fields(input)

		if len(command) == 0 {
			continue
		}

		switch strings.ToUpper(command[0]) {
		case "IN":
			fmt.Println("case in")
		case "OUT":
			fmt.Println("case OUT")
		case "ROUNDROBIN":
			fmt.Println("case ROUNDROBIN")
		case "DEFAULT":
			fmt.Println("case DEFAULT")
		case "EXIT":
			fmt.Println("case EXIT")
			return
		default:
			fmt.Println("case default")
		}
	}

}
