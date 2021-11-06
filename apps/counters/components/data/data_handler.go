package data

import (
	"math"
)

const (
	PLIMIT = 50
	FLIMIT = 50

	ONE = 1

	WHOLES   = "wholes"
	NATURALS = "naturals"
	PRIMES   = "primes"
	FIBS     = "fibs"
)

type (
	Generator struct {
		ActiveSeq map[string]bool
		Displayed uint64
		Index     int
		Cache     map[string][]uint64
		Step      uint64
		Start     uint64
	}
)

func (gen Generator) GetActiveSequence() string {
	var activeSeq string
	for k := range gen.ActiveSeq {
		if gen.ActiveSeq[k] {
			activeSeq = k
		}
	}
	return activeSeq
}

func (gen *Generator) SetActiveSequence(active string) {
	for k := range gen.ActiveSeq {
		if k == active {
			gen.ActiveSeq[k] = true
		} else {
			gen.ActiveSeq[k] = false
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
func (gen *Generator) GenPrimes(length int) {
	if len(gen.Cache[PRIMES]) == 0 {
		gen.Cache[PRIMES] = make([]uint64, length)
		gen.Cache[PRIMES][0] = 2
		num := uint64(3)
		i := 1

		for i < length {
			if isPrime(num) {
				gen.Cache[PRIMES][i] = num
				i++
			}
			num += 2
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

func GetFibByIndex(n uint64) uint64 {
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
func (gen *Generator) GenFibs(length int) {
	if len(gen.Cache[FIBS]) == 0 {
		gen.Cache[FIBS] = make([]uint64, length)
		index := uint64(0)
		for i := range gen.Cache[FIBS] {
			gen.Cache[FIBS][i] = GetFibByIndex(index)
			index++
		}
	}
}

var CurrVals = &Generator{
	Displayed: ONE,
	Step:      ONE,
	Start:     ONE,
	ActiveSeq: map[string]bool{
		WHOLES:   true,
		NATURALS: false,
		PRIMES:   false,
		FIBS:     false,
	},
	Cache: make(map[string][]uint64),
}
