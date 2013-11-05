package sm

import (
	"fmt"
	"testing"
)

const (
	state1 = iota
	state2
	state3
	state4
	state5
	state6
)

var (
	State1 = &State{
		Value:      state1,
		NextStates: []int{state2, state3},
	}
	State2 = &State{
		Value:      state2,
		NextStates: []int{state3},
	}
	State3 = &State{
		Value:      state3,
		NextStates: []int{state5},
	}
	State4 = &State{
		Value:      state4,
		NextStates: []int{state6},
	}
	State5 = &State{
		Value:      state5,
		NextStates: []int{state6},
	}
	State6 = &State{
		Value:      state6,
		NextStates: []int{},
	}
)

func TestValidTransitions(t *testing.T) {
	sm := NewStateMachine(State1)
	if sm.Current() != State1.Value {
		fmt.Println(ErrInitialisationFailure.Error())
		t.FailNow()
	}

	transitions := []*State{State2, State3, State5, State6}
	for _, s := range transitions {
		if !sm.Transition(s) {
			fmt.Println(ErrTransitionFailure.Error())
			t.FailNow()
		} else if sm.Current() != s.Value {
			fmt.Printf(ErrTransitionFailure.Error())
			t.FailNow()
		}
	}

	if !sm.End() {
		fmt.Println(ErrTransitionFailure.Error())
		t.FailNow()
	}
}

func TestInvalidTransitions(t *testing.T) {
	sm := NewStateMachine(State1)
	if sm.Current() != State1.Value {
		fmt.Println(ErrInitialisationFailure.Error())
		t.FailNow()
	}

	transitions := []*State{State2, State3}
	for _, s := range transitions {
		if !sm.Transition(s) {
			fmt.Println(ErrTransitionFailure.Error())
			t.FailNow()
		} else if sm.Current() != s.Value {
			fmt.Printf(ErrTransitionFailure.Error())
			t.FailNow()
		}
	}

	if sm.Transition(State4) || !sm.End() {
		fmt.Println(ErrTransitionFailure.Error())
		t.FailNow()
	}
}
