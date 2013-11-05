package sm

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	exExState1 = 1
	exExState2 = 2
	exExState3 = 3
	exExState4 = 4
	exExState5 = 5
	exExState6 = 6
)

// ReadPrompt prints the prompt (with no modification), and reads
// a line of input from the user. This line is returned as an
// integer.
func ReadPrompt(prompt string) (v int, err error) {
	fmt.Printf("%s", prompt)
	rd := bufio.NewReader(os.Stdin)
	line, err := rd.ReadString('\n')
	if err != nil {
		return
	}
	line = strings.TrimSpace(line)
	return strconv.Atoi(line)
}

var (
	ExState1 = &State{
		Value:      exExState1,
		NextStates: []int{exExState2, exExState3},
	}
	ExState2 = &State{
		Value:      exExState2,
		NextStates: []int{exExState2, exExState4},
	}
	ExState3 = &State{
		Value:      exExState3,
		NextStates: []int{exExState5},
	}
	ExState4 = &State{
		Value:      exExState4,
		NextStates: []int{exExState6},
	}
	ExState5 = &State{
		Value:      exExState5,
		NextStates: []int{exExState6},
	}
	ExState6 = &State{
		Value:      exExState6,
		NextStates: []int{},
	}
)

var States = []*State{nil, ExState1, ExState2, ExState3, ExState4, ExState5, ExState6}

// This is an example of an interactive state machine.
func main() {
	statem := NewStateMachine(ExState1)
	for {
		if statem.Failed() {
			fmt.Println("state machine is in an invalid state: aborting.")
			break
		} else if statem.End() {
			fmt.Println("State machine has finished.")
			break
		}
		fmt.Printf("Current state: State%d\n", statem.Current())
		fmt.Println("Valid states are 1-6. Enter -1 to stop.")
		fmt.Printf("Valid transitions are ")
		for _, st := range States[statem.Current()].NextStates {
			fmt.Printf("%d, ", st)
		}
		fmt.Printf("\n")
		newState, err := ReadPrompt("Next state: ")
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else if newState == -1 {
			break
		} else if newState > 0 && newState < 7 {
			if statem.Transition(States[newState]) {
				fmt.Println("State machine transitioned successfully.")
			} else {
				fmt.Println("State machine failed to transition to", States[newState].Value)
			}
		} else {
			fmt.Println("Invalid state.")
		}
	}
}
