package router

import (
	"fmt"
	"github.com/ovadiaK/bingo/router/handler"
	"net/http"
	"os"
)

var assetsDir = os.Getenv("ASSETS_DIR")
var imagesDir = os.Getenv("IMAGES_DIR")

func New() *http.ServeMux {
	r := http.NewServeMux()
	fmt.Println("bingo-server running, go to")
	r.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(assetsDir))))
	r.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(imagesDir))))
	r.HandleFunc("/game/", handler.Game)
	r.HandleFunc("/cards/", handler.Cards)
	r.HandleFunc("/", handler.Index)
	return r
}
