package utils

import (
	"math"
)

type CurrentValues struct {
	Enabled    bool
	CurrVal    string
	Count      int64
	UCount     uint64
	CountUnit  int64
	UCountUnit uint64
	ResetVal   int64
	UResetVal  uint64
	Wholes
	Naturals
	Primes
	Fibs
}

type Primes struct {
	PEnabled   bool
	PCurrIndex int
	PCache     []uint64
}

type Fibs struct {
	FEnabled   bool
	FCurrIndex int
	FCache     []int64
}

type Wholes struct {
	WEnabled bool
}

type Naturals struct {
	NEnabled bool
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

	countCached := len(p.PCache)
	if length <= countCached {
		return
	} else {
		diff := length - countCached
		startVal := countCached - 1
		lastCached := p.PCache[startVal]
		for diff >= 0 {
			if isPrime(lastCached) {
				p.PCache[startVal+1] = lastCached
			}
			lastCached++
			diff--
		}
	}
}

func (f *Fibs) GetFibByIndex(n uint) int64 {
	if n <= 1 {
		return int64(n)
	}
	var n2, n1 int64 = 0, 1

	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}
	return n2 + n1
}

// GenFibs - Generate Fibonacci sequence
func (f *Fibs) GenFibs(length int) {
	if len(f.FCache) == 0 {
		index := uint(0)
		for len(f.FCache) < length {
			f.FCache = append(f.FCache, f.GetFibByIndex(index))
			index++
		}
	}
}
