package data

import (
	"math"
)

const (
	WHOLES   = "wholes"
	NATURALS = "naturals"
	PRIMES   = "primes"
	FIBS     = "fibs"
)

type CurrentValues struct {
	Enabled        bool
	ActiveSequence map[string]bool
	Primes
	Fibs
	Naturals
	Wholes
}

type Primes struct {
	PCount     uint64
	PCountUnit uint64
	PResetVal  uint64
	PCurrIndex int
	PCache     []uint64
}

type Fibs struct {
	FCount     uint64
	FCountUnit uint64
	FResetVal  uint64
	FCurrIndex int
	FCache     []uint64
}

type Naturals struct {
	NCount     uint64
	NCountUnit uint64
	NResetVal  uint64
}

type Wholes struct {
	WCount     int64
	WCountUnit int64
	WResetVal  int64
}

func (cv CurrentValues) GetActiveSequence() string {
	var activeSeq string
	for k := range CounterVals.ActiveSequence {
		if CounterVals.ActiveSequence[k] {
			activeSeq = k
		}
	}
	return activeSeq
}

func (cv CurrentValues) SetActiveSequence(active string) {
	for k := range CounterVals.ActiveSequence {
		if k == active {
			CounterVals.ActiveSequence[k] = true
		} else {
			CounterVals.ActiveSequence[k] = false
		}
	}
}

func isPrime(n uint64) bool {
	if (n%2 == 0 && n != 2) || (n%3 == 0 && n != 3) || (n%5 == 0 && n != 5) {
		return false
	}

	sqRoot := uint64(math.Sqrt(float64(n)))
	for i := uint64(2); i <= sqRoot; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// GenPrimes - Generate Prime sequence
func (p *Primes) GenPrimes(length int) {
	if len(p.PCache) == 0 {
		num := uint64(2)
		for len(p.PCache) < length {
			if isPrime(num) {
				p.PCache = append(p.PCache, num)
			}
			num++
		}
	}

	//TODO: currently not working
	//countCached := len(p.PCache)
	//if length <= countCached {
	//	return
	//} else {
	//	diff := length - countCached
	//	startVal := countCached - 1
	//	lastCached := p.PCache[startVal]
	//	for diff >= 0 {
	//		if isPrime(lastCached) {
	//			p.PCache[startVal+1] = lastCached
	//		}
	//		lastCached++
	//		diff--
	//	}
	//}
}

func (f *Fibs) GetFibByIndex(n uint64) uint64 {
	if n <= 1 {
		return n
	}
	var n2, n1 uint64 = 0, 1
	for i := uint64(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}
	return n2 + n1
}

// GenFibs - Generate Fibonacci sequence
func (f *Fibs) GenFibs(length int) {
	if len(f.FCache) == 0 {
		index := uint64(0)
		for len(f.FCache) < length {
			f.FCache = append(f.FCache, f.GetFibByIndex(index))
			index++
		}
	}
}

var CounterVals = &CurrentValues{
	Enabled: true,
	ActiveSequence: map[string]bool{
		WHOLES:   true,
		NATURALS: false,
		PRIMES:   false,
		FIBS:     false,
	},
	Primes: Primes{
		PCount:     uint64(2),
		PCountUnit: uint64(1),
		PResetVal:  uint64(0),
	},
	Fibs: Fibs{
		FCount:     uint64(0),
		FCountUnit: uint64(1),
		FResetVal:  uint64(0),
	},
	Naturals: Naturals{
		NCount:     uint64(0),
		NCountUnit: uint64(1),
		NResetVal:  uint64(0),
	},
	Wholes: Wholes{
		WCount:     int64(0),
		WCountUnit: int64(1),
		WResetVal:  int64(0),
	},
}
