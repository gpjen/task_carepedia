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
	nextGender string
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

	patientOut := ""

	if h.Order == "default" {
		patientOut = fmt.Sprintf("%s %s", h.Queue[0].MRnumber, h.Queue[0].Gender)
		h.Queue = h.Queue[1:]

	} else if h.Order == "roudRobin" {
		for i, patient := range h.Queue {
			if patient.Gender == h.nextGender {
				nextGender := "M"
				if h.nextGender == "M" {
					nextGender = "F"
				}
				h.nextGender = nextGender
				patientOut = fmt.Sprintf("%s %s", patient.MRnumber, patient.Gender)
				h.Queue = append(h.Queue[:i], h.Queue[i+1:]...)
				break
			}
		}
		if patientOut == "" {
			patientOut = fmt.Sprintf("%s %s", h.Queue[0].MRnumber, h.Queue[0].Gender)
			h.Queue = h.Queue[1:]
		}
	}

	fmt.Printf("send : %s\n", patientOut)
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
		nextGender: "F",
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
			return
		default:
			fmt.Println("Invalid command.")
		}
	}

}
