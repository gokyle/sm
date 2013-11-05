// state implements a simple finite state machine. It is designed
// for keep track of valid states and transitions for simple systems.
package sm

import "errors"

// State represents a single state and the possible states that it
// can proceed to.
type State struct {
	Value      int
	NextStates []int
}

var (
	ErrTransitionFailure     = errors.New("state: invalid transition")
	ErrInitialisationFailure = errors.New("state: failed to initialise state machine")
)

// InvalidState is a predefined state indicating that the state
// machine has entered into an inconsistent state. This state is
// used to represent exceptional conditions or critical errors from
// which the system should not recover.
var InvalidState = &State{
	Value:      -1,
	NextStates: []int{},
}

// StateMachine is a representation of a simple finite state machine.
// The state machine is initialised to a particular state, and
// transitions along a set sequence of states until it reaches a
// terminating state.
type StateMachine struct {
	State *State
}

// End returns true if the state machine is at a terminating
// condition.
func (state *StateMachine) End() bool {
	if len(state.State.NextStates) == 0 {
		return true
	} else if state.State == InvalidState {
		return true
	}
	return false
}

// Current returns the current state.
func (state *StateMachine) Current() int {
	return state.State.Value
}

// Fail places the machine into an inconsistent state.
func (state *StateMachine) Fail() {
	state.State = InvalidState
}

// Failed returns true if the state machine is in an invalid state.
func (state *StateMachine) Failed() bool {
	return state.State == InvalidState
}

// Transition attempts to move the state machine to the given state.
// If the state being transitioned to isn't a valid state that the
// current state is allowed to transition to, the state machine is
// placed into the InvalidState.
func (state *StateMachine) Transition(to *State) bool {
	for _, ns := range state.State.NextStates {
		if ns == to.Value {
			state.State = to
			return true
		}
	}
	state.State = InvalidState
	return false
}

// NewStateMachine creates a new state machine starting at the given
// state.
func NewStateMachine(initial *State) (state *StateMachine) {
	return &StateMachine{State: initial}
}
