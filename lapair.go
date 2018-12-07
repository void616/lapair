package lapair

// Gender type
type Gender uint8

const (
	// Masculine gender
	Masculine Gender = 0
	// Feminine gender
	Feminine Gender = 1
	// Neuter gender
	Neuter Gender = 2
	// Genders count
	genders uint8 = 3
)

// Pair getter
func Pair(n, a uint16) (string, string) {
	nou, gen := Noun(n)
	adj := Adjective(a, gen)
	return nou, adj
}

// Noun getter
func Noun(n uint16) (string, Gender) {
	ret := latinN[int(n)%len(latinN)]
	return ret.v, ret.g
}

// Adjective getter
func Adjective(a uint16, g Gender) string {
	ret := latinA[int(a)%len(latinA)]
	switch len(ret) {
	case 3:
		return ret[uint8(g)%genders]
	case 2:
		switch g {
		case Masculine, Feminine:
			return ret[0]
		default:
			return ret[1]
		}
	default:
		return ret[0]
	}
}
