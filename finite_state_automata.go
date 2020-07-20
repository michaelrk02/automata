package automata

import (
    "fmt"
)

type FSATransition = func(token byte) int

type FiniteStateAutomata struct {
    Automata

    States []FSATransition

    start, final, cur int
}

func NewFiniteStateAutomata(nstates, start, final int) *FiniteStateAutomata {
    fsa := new(FiniteStateAutomata)
    fsa.States = make([]FSATransition, nstates)
    fsa.start = start
    fsa.final = final
    return fsa
}

func (fsa *FiniteStateAutomata) Accepts(in string) bool {
    fsa.cur = fsa.start

    for i := 0; i < len(in); i++ {
        fsa.feed(in[i])
    }
    fsa.feed(0x00)

    return fsa.cur == fsa.final
}

func (fsa *FiniteStateAutomata) feed(token byte) {
    fn := fsa.States[fsa.cur]
    if fn == nil {
        panic(fmt.Sprintf("no transition function is available for input %v on state %d\n", token, fsa.cur))
    }

    fsa.cur = fn(token)
}

