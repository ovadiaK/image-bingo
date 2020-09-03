package collection

import (
	"fmt"
)

type Matrix []row

type row []Image

//const rowLength int = 3
const colLen int = 10
const hiddenPath string = "../images/hidden/question.png"

func (s Set) Matrix() (Matrix, error) {
	if len(s.Symbols) == 0 {
		return nil, fmt.Errorf("no logo initiated")
	}
	rowLength := len(s.Symbols) / colLen
	if len(s.Symbols)%colLen != 0 {
		rowLength++
	}
	mtx := Matrix{}
	d := 1
	newImage := Image{}
	for i := 0; i < colLen; i++ {
		r := row{}
		for j := 0; j < rowLength; j++ {
			newImage.Number = d
			newImage.Path = hiddenPath
			newImage.Title = "unrevealed"
			r = append(r, newImage)
			d++
		}
		mtx = append(mtx, r)
	}
	return mtx, nil

}
