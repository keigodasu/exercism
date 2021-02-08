package dominoes

type Domino [2]int

func MakeChain(dominos []Domino) (chain []Domino, ok bool)  {
	if l := len(dominos); l == 0 {
		ok = true
		return
	} else if l == 1 {
		if ok = dominos[0][0] == dominos[0][1]; ok {
			chain = dominos
		}
		return
	}

	var (
		dominoFlags []bool
		makeChain func(int) bool
	)

	dominoFlags = make([]bool, len(dominos))
	makeChain = func(n int) bool {
		if n == len(dominos) {
			return true
		}

		for i, domino := range dominos {
			if !dominoFlags[i] {
				switch chain[n-1][1] {
				case domino[0]:
					chain = append(chain, domino)
				case domino[1]:
					chain = append(chain, Domino{domino[1], domino[0]})
				default:
					continue
				}
				dominoFlags[i] = true
				if makeChain(n + 1) {
					return true
				}
				dominoFlags[i] = false
				chain = chain[:n]
			}
		}
		return false
	}

	for i, domino := range dominos {
		chain, dominoFlags[i] = []Domino{domino}, true
		if ok = makeChain(1); ok {
			break
		}
	}

	return
}
