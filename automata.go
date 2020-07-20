package automata

type Automata interface {
    Accepts(in string) bool
}

