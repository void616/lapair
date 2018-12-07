package lapair

import (
	"regexp"
	"testing"
)

func TestArrayAdjective(t *testing.T) {
	if len(latinA) == 0 || len(latinA) > 0xFFFF {
		t.Errorf("Invalid length: %v", len(latinA))
	}
	for i, v := range latinA {
		if len(v) < 1 || len(v) > 3 {
			t.Errorf("Invalid adjective array length at %v", i)
		}
	}
	rex := regexp.MustCompile("^[a-z]+$")
	for i, v := range latinA {
		for j, a := range v {
			if !rex.MatchString(a) {
				t.Errorf("Invalid item at [%v][%v] (%v)", i, j, a)
			}
		}
	}
}

func TestArrayNoun(t *testing.T) {
	if len(latinN) == 0 || len(latinN) > 0xFFFF {
		t.Errorf("Invalid length: %v", len(latinN))
	}
	rex := regexp.MustCompile("^[a-z]+$")
	for i, v := range latinN {
		if !rex.MatchString(v.v) {
			t.Errorf("Invalid item at [%v] (%v)", i, v.v)
		}
	}
}
