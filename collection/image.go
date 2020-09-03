package collection

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Set struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Kind    string
	Icon    Image
	Symbols []Image
}

type Image struct {
	Path   string `json:"path"`
	Number int    `json:"number"`
	Title  string `json:"title"`
}

type ImageFile struct {
	Resolution int    `json:"resolution"`
	Path       string `json:"path"`
}

var id int64 = 1
var imageExtensions = []string{".jpg", ".png", ".JPG"}

const minSymbols = 20

func Read(dirName string) (*Set, error) {
	fi, err := ioutil.ReadDir(filepath.Join(dirName))
	if err != nil {
		return nil, err
	}
	res := Set{}
	res.Name = name(dirName)
	for _, f := range fi {
		if strings.TrimSuffix(f.Name(), path.Ext(f.Name())) == "icon" {
			res.Icon.Path = pathCorrect(filepath.Join(dirName, f.Name()))
		}
		if f.IsDir() && f.Name() == "pics" {
			symbols, err := readImages(filepath.Join(dirName, f.Name()))
			if err != nil {
				return nil, err
			}
			res.Symbols = symbols
		}
	}
	if len(res.Symbols) < minSymbols {
		return nil, fmt.Errorf("collection %v too few symbols", res.Name)
	}
	if res.Icon.Path == "" {
		res.Icon.Path = res.Symbols[0].Path
	}
	res.Id = id
	id++
	return &res, nil
}

func ReadAll(dirName string, exclude ...string) ([]Set, []error, error) {
	errs := make([]error, 0)
	infos, err := ioutil.ReadDir(dirName)
	if err != nil {
		return nil, errs, err
	}
	resSet := make([]Set, 0, len(infos))
	for _, info := range infos {
		if info.IsDir() {
			if !stringInSlice(info.Name(), exclude) {
				s, err := Read(filepath.Join(dirName, info.Name()))
				if err != nil {
					errs = append(errs, err)
				} else {
					resSet = append(resSet, *s)
				}
			}
		}
	}
	if len(resSet) == 0 {
		return resSet, errs, fmt.Errorf("no sets found")
	}
	return resSet, errs, err
}

// orders the Set according to given order
func (s *Set) Order(order []int) {
	newOrder := make([]Image, len((*s).Symbols))
	for i, j := range order {
		newOrder[i] = (*s).Symbols[j]
	}
	(*s).Symbols = newOrder
}

func (i *Image) CutTitle(n int) {
	if len(i.Title) > n {
		i.Title = i.Title[:n]

	}
}

func readImages(dirName string) ([]Image, error) {
	fi, err := ioutil.ReadDir(dirName)
	if err != nil {
		return nil, err
	}
	res := make([]Image, 0, len(fi))
	i := 1
	for _, f := range fi {
		if stringInSlice(filepath.Ext(f.Name()), imageExtensions) {
			img := Image{
				Path:   pathCorrect(path.Join(dirName, f.Name())),
				Number: i,
				Title:  filepath.Base(strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))),
			}
			res = append(res, img)
			i++
		}
	}
	return res, nil
}

// return name of collection
func (s *Set) name() string {
	sep := string(os.PathSeparator)
	for _, i := range s.Symbols {
		dir, _ := filepath.Split(i.Path)
		dirs := strings.Split(strings.Trim(dir, sep), sep)
		return dirs[len(dirs)-2]
	}
	return ""
}
func name(dirName string) string {
	sep := string(os.PathSeparator)
	dirs := strings.Split(strings.Trim(dirName, sep), sep)
	if len(dirs) > 1 {
		return dirs[len(dirs)-1]
	}
	return ""
}
func pathCorrect(in string) string {
	return path.Join("../images", strings.TrimPrefix(in, os.Getenv("IMAGES_DIR")))
}
