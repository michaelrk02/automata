package automata

import (
    "fmt"
)

type TMHeadDir = int
const (
    TMHeadDirLeft TMHeadDir = iota
    TMHeadDirRight
)

type TMTransition = func(token byte) (int, byte, TMHeadDir)

type TuringMachine struct {
    Automata

    States []TMTransition

    start, accept, reject, cur int
}

func NewTuringMachine(nstates, start, accept, reject int) *TuringMachine {
    tm := new(TuringMachine)
    tm.States = make([]TMTransition, nstates)
    tm.start = start
    tm.accept = accept
    tm.reject = reject
    return tm
}

func (tm *TuringMachine) Accepts(in string) bool {
    arr := make([]byte, len(in) + 1)
    copy(arr, in)

    tm.cur = tm.start

    i := 0
    for true {
        repl, dir := tm.feed(arr[i])
        arr[i] = repl
        switch dir {
        case TMHeadDirLeft:
            i = (i - 1) % len(arr)
        case TMHeadDirRight:
            i = (i + 1) % len(arr)
        }
        if tm.cur == tm.accept {
            return true
        }
        if tm.cur == tm.reject {
            break
        }
    }

    return false
}

func (tm *TuringMachine) feed(token byte) (byte, TMHeadDir) {
    fn := tm.States[tm.cur]
    if fn == nil {
        panic(fmt.Sprintf("no transition function is available for input %v on state %d\n", token, tm.cur))
    }

    var repl byte
    var dir TMHeadDir
    tm.cur, repl, dir = fn(token)

    return repl, dir
}
