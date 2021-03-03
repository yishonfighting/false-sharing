package False_Sharing

import (
	"sync"
	"testing"
)

func testAtomicIncrease(atm MyAtomic) {
	paraNum := 1000
	addTimes := 1000
	var wg sync.WaitGroup
	wg.Add(paraNum)
	for i := 0; i < paraNum; i++ {
		go func() {
			for j := 0; j < addTimes; j++ {
				atm.IncreaseAllEles()
			}
			wg.Done()
		}()
	}
	wg.Wait()

}
func BenchmarkNoPad(b *testing.B) {
	atm := &NoPad{}
	b.ResetTimer()
	testAtomicIncrease(atm)
}

func BenchmarkPad(b *testing.B) {
	atm := &Pad{}
	b.ResetTimer()
	testAtomicIncrease(atm)
}