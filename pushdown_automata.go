package automata

import (
    "fmt"
)

type PDAStackOp interface{}

type PDAStackOpPush struct {
    PDAStackOp

    Symbol byte
}

type PDAStackOpPop struct {
    PDAStackOp
}

type PDAStackOpIgnore struct {
    PDAStackOp
}

type PDATransition = func(token, top byte) (int, PDAStackOp)

type PushDownAutomata struct {
    Automata

    States []PDATransition

    start, final, cur int
    stack []byte
}

func NewPushDownAutomata(nstates, start, final int) *PushDownAutomata {
    pda := new(PushDownAutomata)
    pda.States = make([]PDATransition, nstates)
    pda.start = start
    pda.final = final
    pda.stack = make([]byte, 0)
    return pda
}

func (pda *PushDownAutomata) Accepts(in string) bool {
    pda.cur = pda.start
    for i := 0; i < len(in); i++ {
        pda.feed(in[i])
    }
    pda.feed(0x00)

    return pda.cur == pda.final
}

func (pda *PushDownAutomata) feed(token byte) {
    fn := pda.States[pda.cur]
    if fn == nil {
        panic(fmt.Sprintf("no transition function is available for input %v on state %d\n", token, pda.cur))
    }

    var op PDAStackOp
    pda.cur, op = fn(token, pda.top())

    switch op.(type) {
    case PDAStackOpPush:
        pda.push(op.(*PDAStackOpPush).Symbol)
    case PDAStackOpPop:
        pda.pop()
    case PDAStackOpIgnore:
    }
}

func (pda *PushDownAutomata) top() byte {
    return pda.stack[len(pda.stack) - 1]
}

func (pda *PushDownAutomata) push(symbol byte) {
    pda.stack = append(pda.stack, symbol)
}

func (pda *PushDownAutomata) pop() {
    pda.stack = pda.stack[0:len(pda.stack) - 1]
}

