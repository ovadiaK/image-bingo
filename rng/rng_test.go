package rng

import (
	"fmt"
	"sort"
	"testing"
)

const (
	ARRAY_LENGTH = 30
	SUBSET       = 9
	AMOUNT       = 20
)

func TestArray1toN(t *testing.T) {
	a := array1toN(ARRAY_LENGTH)
	if len(a) != ARRAY_LENGTH {
		t.Fatal("len not correct")
	}
	for i := 0; i < ARRAY_LENGTH; i++ {
		if a[i] != i {
			t.Fatal(fmt.Sprint(a[i], " != ", i))
		}
	}
}
func TestInitiate(t *testing.T) {
	a := Initiate(ARRAY_LENGTH)
	if len(a) != ARRAY_LENGTH {
		t.Fatal("len not correct")

	}
	sort.Ints(a)
	for i := 0; i < ARRAY_LENGTH; i++ {
		if a[i] != i {
			t.Fatal(fmt.Sprint(a[i], " != ", i))
		}
	}
}

func TestStack(t *testing.T) {
	s := Stack(ARRAY_LENGTH, SUBSET, AMOUNT)
	fmt.Println(s)
	if len(s) != AMOUNT {
		t.Fatal("amount is %i,expected %i", len(s), AMOUNT)
	}

}

//todo selection and id
// length of subset
// id correct
