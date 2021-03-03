package False_Sharing

import "sync/atomic"

type MyAtomic interface {
	Increase()
}

type NoPad struct {
	a uint64
	b uint64
	c uint64
}

func (atm *NoPad) Increase() {
	atomic.AddUint64(&atm.a,1)
	atomic.AddUint64(&atm.b,1)
	atomic.AddUint64(&atm.c,1)
}

type Pad struct {
	a uint64
	_p1 [8]uint64

	b uint64
	_p2 [8]uint64

	c uint64
	_p3 [8]uint64
}

func (atm *Pad) Increase() {
	atomic.AddUint64(&atm.a,1)
	atomic.AddUint64(&atm.b,1)
	atomic.AddUint64(&atm.c,1)
}
