package handler

import (
	"github.com/ovadiaK/bingo/collection"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type messageCards struct {
	Cards []cardPair
}
type cardPair struct {
	CardA collection.Card
	CardB collection.Card
}

func Cards(w http.ResponseWriter, r *http.Request) {
	dir := r.FormValue("collection")
	amount := r.FormValue("amount")
	var a int
	if len(amount) < 9 {
		var err error
		a, err = strconv.Atoi(amount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "To big input", 422)
	}
	if a == 0 {
		a = 4
	}
	if a < 0 {
		a = -a
	}
	if a > 40 {
		a = 40
	}
	if a%4 != 0 {
		a += 4 - (a % 4)
	}

	mess, err := newCards(filepath.Join(imageFolder, dir), a, picsPerCards)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	renderTemplate(w, "cards", mess)
}
func newCards(dirName string, amount, subset int) (messageCards, error) {
	mess := messageCards{}
	imgs, err := collection.Read(filepath.Join(dirName))
	if err != nil {
		return mess, err
	}
	if amount%2 == 0 {
		amount++
	}
	mess.Cards = make([]cardPair, 0, amount/2)
	cards, err := imgs.Stack(subset, amount)
	card := cardPair{}
	for i, c := range cards {
		if (i+1)%2 == 0 {
			card.CardB = c
			mess.Cards = append(mess.Cards, card)
			continue
		}
		card.CardA = c
	}
	return mess, err
}
