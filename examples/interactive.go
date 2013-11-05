package main

import (
	"bufio"
	"fmt"
	"github.com/gokyle/state"
	"os"
	"strconv"
	"strings"
)

const (
	state1 = 1
	state2 = 2
	state3 = 3
	state4 = 4
	state5 = 5
	state6 = 6
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
	State1 = &sm.State{
		Value:      state1,
		NextStates: []int{state2, state3},
	}
	State2 = &sm.State{
		Value:      state2,
		NextStates: []int{state2, state4},
	}
	State3 = &sm.State{
		Value:      state3,
		NextStates: []int{state5},
	}
	State4 = &sm.State{
		Value:      state4,
		NextStates: []int{state6},
	}
	State5 = &sm.State{
		Value:      state5,
		NextStates: []int{state6},
	}
	State6 = &sm.State{
		Value:      state6,
		NextStates: []int{},
	}
)

var States = []*sm.State{nil, State1, State2, State3, State4, State5, State6}

// This is an example of an interactive state machine.
func main() {
	sm := sm.NewStateMachine(State1)
	for {
		if sm.Failed() {
			fmt.Println("state machine is in an invalid state: aborting.")
			break
		} else if sm.End() {
			fmt.Println("State machine has finished.")
			break
		}
		fmt.Printf("Current state: State%d\n", sm.Current())
		fmt.Println("Valid states are 1-6. Enter -1 to stop.")
		fmt.Printf("Valid transitions are ")
		for _, st := range States[sm.Current()].NextStates {
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
			if sm.Transition(States[newState]) {
				fmt.Println("State machine transitioned successfully.")
			} else {
				fmt.Println("State machine failed to transition to", States[newState].Value)
			}
		} else {
			fmt.Println("Invalid state.")
		}
	}
}
