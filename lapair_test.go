package lapair

import "testing"

func TestPair(t *testing.T) {

	t.Log(Pair(0, 0))
	t.Log(Pair(0xFF, 0xFF))
	t.Log(Pair(0xFFFF, 0xFFFF))
}

func TestAdjective(t *testing.T) {

	Adjective(0, Masculine)
	Adjective(0xF, Feminine)
	Adjective(0xFF, Neuter)
	Adjective(0xFFF, Masculine)
	Adjective(0xFFFF, Feminine)
	Adjective(0xFFFF, Gender(Neuter+1))
	Adjective(0x1337, Gender(Neuter+2))
	Adjective(0xBEEF, Gender(Neuter+3))

	for i, v := range latinA {
		if len(v) == 1 {
			Adjective(uint16(i), Masculine)
			Adjective(uint16(i), Feminine)
			Adjective(uint16(i), Neuter)
			break
		}
	}

	for i, v := range latinA {
		if len(v) == 2 {
			Adjective(uint16(i), Masculine)
			Adjective(uint16(i), Feminine)
			Adjective(uint16(i), Neuter)
			break
		}
	}

	for i, v := range latinA {
		if len(v) == 3 {
			Adjective(uint16(i), Masculine)
			Adjective(uint16(i), Feminine)
			Adjective(uint16(i), Neuter)
			break
		}
	}
}

func TestNoun(t *testing.T) {

	Noun(0)
	Noun(0xF)
	Noun(0xFF)
	Noun(0xFFF)
	Noun(0xFFFF)
}
