package handler

import (
	"fmt"
	"github.com/ovadiaK/bingo/collection"
	"log"
	"net/http"
)

type messageSetup struct {
	Collections []setInfo `json:"collections"`
	Command     startInfo `json:"command"`
}

type startInfo struct {
	Request string
	Name    string
	Button  string
}
type setInfo struct {
	Name       string `json:"name"`
	Size       int    `json:"size"`
	TitleImage string `json:"title_image"`
	Card       collection.Card
}

func Index(w http.ResponseWriter, r *http.Request) {
	sets, _, err := collection.ReadAll(imageFolder, hiddenFolder)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	setInfos := make([]setInfo, 0, len(sets))
	for _, s := range sets {
		card, err := s.Stack(9, 1)
		if err != nil {
			fmt.Println(err)
			continue
		}
		info := setInfo{
			Name:       s.Name,
			Size:       len(s.Symbols),
			TitleImage: s.Icon.Path,
			Card:       card[0],
		}
		setInfos = append(setInfos, info)
	}
	renderTemplate(w, "index", messageSetup{
		Collections: setInfos,
	})
}
