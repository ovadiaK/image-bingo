package collection

import (
	"fmt"
	"github.com/ovadiaK/bingo/rng"
)

type Card struct {
	Name string
	Id   string
	Img  [][]Image
}

func (s Set) Stack(subset, amount int) ([]Card, error) {
	if len(s.Symbols) == 0 {
		return nil, fmt.Errorf("no logo initiated")
	}
	if subset <= 0 {
		return nil, fmt.Errorf("incorrect number of logo per card: %v", subset)
	}
	if amount <= 0 {
		return nil, fmt.Errorf("incorrect stack size requested: %v", amount)
	}
	cards := make([]Card, 0, len(s.Symbols))
	rmap := rng.Stack(len(s.Symbols), subset, amount)
	for id, nums := range rmap {
		c := Card{Id: id}
		c.Name = s.Name
		row := make([]Image, 3)
		for j, n := range nums {
			s.Symbols[n].CutTitle(13)
			row[j%3] = s.Symbols[n]
			if j != 0 && (j+1)%3 == 0 {
				c.Img = append(c.Img, row)
				row = make([]Image, 3)
			}
		}
		cards = append(cards, c)
	}

	return cards, nil
}
