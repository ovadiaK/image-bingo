package rng

import (
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func Initiate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	a := array1toN(n)
	res := make([]int, 0, n)
	m := n
	for i := 0; i < n; i++ {
		b := rand.Intn(m)
		res = append(res, a[b])
		a = append(a[:b], a[b+1:]...)
		m--
	}
	return res
}
func Stack(n, subset, amount int) map[string][]int {
	stack := make(map[string][]int)
	j := 0
	for j < amount {
		r := selection(n, subset)
		sort.Ints(r)
		i := id(n, r)
		if _, ok := stack[i]; !ok {
			stack[i] = r
			j++
		}
	}
	return stack
}
func selection(n, subset int) []int {
	r := Initiate(n)
	return r[:subset]
}
func array1toN(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}
	return a
}
func id(n int, row []int) string {
	size := len(strconv.Itoa(n))
	var id string
	for _, r := range row {
		str := strconv.Itoa(r)
		for len(str) < size {
			str = "0" + str
		}
		id += str
	}
	return id
}
