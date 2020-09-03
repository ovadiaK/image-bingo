package handler

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	hiddenFolder = "hidden"
	picsPerCards = 9
)

var imageFolder = os.Getenv("IMAGES_DIR")

var (
	templates    *template.Template
	templatesDir = os.Getenv("TEMPLATES_DIR")
)

func init() {
	templates = parseTemplates(templatesDir, ".gohtml")
}
func parseTemplates(path, extension string) *template.Template {
	templ := template.New("")
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, extension) {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}
		return err
	})
	if err != nil {
		panic(err)
	}
	return templ
}

func renderTemplate(w http.ResponseWriter, tmpl string, i interface{}) {
	if err := templates.ExecuteTemplate(w, tmpl+".gohtml", i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
