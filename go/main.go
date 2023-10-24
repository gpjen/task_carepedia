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
	Queue      []Patient
	Order      string
	LastGender string
}

func (h *HospitalQueue) isDuplicateMRNumber(mrNumber string) bool {
	for _, patient := range h.Queue {
		if (patient.MRnumber) == mrNumber {
			return true
		}
	}
	return false
}

func (h *HospitalQueue) HandleIn(patientInput Patient) {
	if patientInput.Gender != "F" && patientInput.Gender != "M" {
		fmt.Println("gender invalid.")
		return
	}

	if h.isDuplicateMRNumber(patientInput.MRnumber) {
		fmt.Printf("error : patient with %s already in queue\n", patientInput.MRnumber)
		return
	}

	h.Queue = append(h.Queue, patientInput)
}

func (h *HospitalQueue) HandleOut() {
	if len(h.Queue) < 1 {
		fmt.Println("queue is empty")
		return
	}

	if h.Order == "default" {
		patient := fmt.Sprintf("%s %s", h.Queue[0].MRnumber, h.Queue[0].Gender)
		fmt.Println(patient)
		h.Queue = h.Queue[1:]

	} else if h.Order == "roudRobin" {
		// for _, v := range v {

		// }
	}
}

func (h *HospitalQueue) HandleRRobin() {
	h.Order = "roudRobin"
}

func (h *HospitalQueue) HandleDefault() {
	h.Order = "default"
}

func main() {
	fmt.Println(" ===[ Hospital Queue Aplication ]===")

	queue := HospitalQueue{
		Queue:      []Patient{},
		Order:      "default",
		LastGender: "F",
	}

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
			if len(command) < 3 {
				fmt.Println("error : invalid IN command")
				break
			}
			queue.HandleIn(Patient{
				MRnumber: fmt.Sprintf("MR%s", command[1]),
				Gender:   strings.ToUpper(command[2]),
			})

			fmt.Println(queue)
		case "OUT":
			queue.HandleOut()
		case "ROUNDROBIN":
			queue.HandleRRobin()
		case "DEFAULT":
			queue.HandleDefault()
		case "EXIT":
			fmt.Println("case EXIT")
			return
		default:
			fmt.Println("Invalid command.")
		}
	}

}
