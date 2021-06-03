package wallet

import "fmt"

type Bitcion int

type Wallet struct {
	bitcion Bitcion
}

func (w *Wallet) Deposit(amount Bitcion) {
	w.bitcion += Bitcion(amount)
}

func (w *Wallet) Balance() Bitcion {
	return w.bitcion
}

type Stringer interface {
	String() string
}

func (b Bitcion) String() string {
	return fmt.Sprintf("%d BTC", b)
}
