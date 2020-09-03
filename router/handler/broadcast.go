package handler

import (
	"encoding/json"
	"github.com/ovadiaK/bingo/collection"
	"github.com/ovadiaK/bingo/rng"
	"log"
	"net/http"
	"path/filepath"
)

type messageBingo struct {
	Matrix collection.Matrix
	Prize  string
	Size   int
	Order  string
	Img    collection.Image
}

func Game(w http.ResponseWriter, r *http.Request) {
	dir := r.FormValue("collection")
	message, err := newGame(filepath.Join(imageFolder, dir))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	message.Prize = r.FormValue("prize")
	renderTemplate(w, "broadcast", message)
}

func newGame(dirName string) (messageBingo, error) {
	mess := messageBingo{}
	s, err := collection.Read(filepath.Join(dirName))
	if err != nil {
		return mess, err
	}
	mess.Size = len(s.Symbols)
	order := rng.Initiate(len(s.Symbols))
	s.Order(order)
	imagesJson, err := json.Marshal(s.Symbols)
	if err != nil {
		return mess, err
	}
	mess.Order = string(imagesJson)
	mess.Matrix, err = s.Matrix()
	if err != nil {
		return mess, err
	}
	mess.Img = mess.Matrix[0][0]
	return mess, err
}
