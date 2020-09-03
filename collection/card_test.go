package collection

import (
	"fmt"
	"testing"
	"time"
)

func TestSet_Stack(t *testing.T) {
	s, err := Read("../images/animals")
	if err != nil {
		fmt.Println(err)
	}
	c, err := s.Stack(9, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(c))
	for _, i := range c {
		fmt.Println(i.Id)
		for _, j := range i.Img {
			fmt.Println(j)
		}
	}
}
func BenchmarkSet_Stack(b *testing.B) {
	s, err := Read("../images/animals")
	if err != nil {
		fmt.Println(err)
	}
	start := time.Now()

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		s.Stack(9, 24000)
	}

	fmt.Println(time.Since(start))
}
