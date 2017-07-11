package projecteular

import (
	"fmt"
	"math"

	"github.com/redisliu/gotest/utils"
)

const cliqueSize = 5

func pe56() {
	M := 10000000
	arr := utils.ComputePrimeArray(M)
	primeList := make([]int64, 0, M)
	primeList = append(primeList, 2)
	primeSet := map[int64]bool{2: true}

	maxL := int64(len(arr)*2 + 1)
	for i := int64(3); i <= maxL; i++ {
		if utils.IsPrime(i, arr) {
			primeList = append(primeList, i)
			primeSet[i] = true
		}
	}
	fmt.Printf("prime: %v\n", primeList[0:100])

	adj := make(map[int64]map[int64]bool)

	for i, r := range primeList {
		wr := int64(math.Pow10(1 + int(math.Log10(float64(r)))))
		for j := 0; j < i; j++ {
			l := primeList[j]
			wl := int64(math.Pow10(1 + int(math.Log10(float64(l)))))
			//fmt.Printf("test for : r: %d, l: %d, ano: %d, ano2: %d\n", r, l, l*wr+r, r*wl+l)
			if utils.IsPrime(l*wr+r, arr) && utils.IsPrime(r*wl+l, arr) {
				if s := addEdge(l, r, adj, cliqueSize); len(s) != 0 {
					fmt.Printf("got result: %v\n graph: \n", s)
					return
				}
			}
		}
	}
}

func addEdge(l, r int64, adj map[int64]map[int64]bool, N int) []int64 {
	// Add edge
	lm, ok := adj[l]
	if ok {
		lm[r] = true
	} else {
		lm = map[int64]bool{r: true}
		adj[l] = lm
	}
	rm, ok := adj[r]
	if ok {
		rm[l] = true
	} else {
		rm = map[int64]bool{l: true}
		adj[r] = rm
	}

	// check
	if len(lm) < N-1 || len(rm) < N-1 {
		return nil
	}

	s := make(map[int64]bool)
	for k := range lm {
		if _, ok := rm[k]; ok {
			s[k] = true
		}
	}
	if len(s) < N-2 {
		return nil
	}

	clique := make([]int64, 0, 5)
	clique = append(clique, l, r)
	for k := range s {
		neighbor := adj[k]
		found := true
		if len(neighbor) < len(clique) {
			found = false
		} else {
			for _, v := range clique {
				if _, ok := neighbor[v]; !ok {
					found = false
					break
				}
			}
		}
		if found {
			clique = append(clique, k)
			if len(clique) == N {
				break
			}
		}
	}
	if len(clique) == N {
		return clique
	}
	// TODO: check clique
	return nil
}
