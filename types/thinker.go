package types

type Think Action

type Thinker struct {
	Prev     *Thinker
	Next     *Thinker
	Function Think
}
